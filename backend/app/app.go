package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
)

type ttsRequest struct {
	Prompt  string `json:"prompt"`
	VoiceId string `json:"voice_id"`
	ModelId string `json:"model_id"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}
	}
}

func tts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ttsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prompt to content
	finalContent, err := GenContent(&req)
	if err != nil {
		http.Error(w, "Failed to generate content: "+err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Final Content:\n", finalContent)

	// text to speech
	resp, err := GenAudio(finalContent, req)
	if err != nil {
		http.Error(w, "Failed to generate audio: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// save audio to blob
	blobURL, err := SaveAudioToBlob(resp.Body)
	if err != nil {
		http.Error(w, "Failed to save audio to blob: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Audio saved to blob: ", blobURL)

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Cache-Control", "no-store")

	// send audio to frontend
	json.NewEncoder(w).Encode(map[string]string{
		"audio_url": blobURL,
	})
}

func SaveAudioToBlob(audio io.Reader) (string, error) {

	if audio == nil {
		return "", fmt.Errorf("audio is nil")
	}

	audioFile, err := io.ReadAll(audio)
	if err != nil {
		return "", err
	}

	VercelBlobSroreId := os.Getenv("BLOB_READ_WRITE_TOKEN")
	fileName := uuid.New().String() + ".mp3"
	blobURL := "https://blob.vercel-storage.com/put"

	req, err := http.NewRequest("PUT", blobURL, bytes.NewReader(audioFile))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "audio/mpeg")
	req.Header.Set("Authorization", "Bearer "+VercelBlobSroreId)
	req.Header.Set("x-vercel-blob-filename", fileName)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	fmt.Println("Response body:", string(body))
	var respData struct {
		URL string `json:"url"`
	}
	if err := json.Unmarshal(body, &respData); err != nil {
		return "", fmt.Errorf("failed to parse blob response: %w", err)
	}
	blobURL = respData.URL

	return blobURL, nil
}

func Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/auth/google/login", handleGoogleLogin)
	http.HandleFunc("/auth/google/callback", handleGoogleCallback)
	http.Handle("/auth/me", AuthMiddleware(http.HandlerFunc(handleProfile)))
	http.Handle("/auth/logout", AuthMiddleware(http.HandlerFunc(handleLogout)))

	http.Handle("/api/tts", AuthMiddleware(http.HandlerFunc(tts)))

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", CorsMiddleware(http.DefaultServeMux))
}
