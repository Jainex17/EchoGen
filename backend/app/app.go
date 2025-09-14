package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"echogen/backend/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type ttsRequest struct {
	Prompt  string `json:"prompt"`
	Style   string `json:"style"`
	VoiceId string `json:"voice_id"`
	ModelId string `json:"model_id"`
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
	resp, err := GenAudio(finalContent, &req)
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

	userRepo := model.UserRepository{}
	audioGenerationRepo := model.AudioGenerationRepository{}

	email := r.Context().Value(ClaimsKey).(jwt.MapClaims)["email"].(string)

	user, err := userRepo.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "Failed to get user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	audioGeneration := &model.AudioGeneration{
		UserID:   user.ID,
		Prompt:   req.Prompt,
		Content:  finalContent,
		AudioURL: blobURL,
	}

	// save content to database
	err = audioGenerationRepo.CreateAudioGeneration(audioGeneration)
	if err != nil {
		http.Error(w, "Failed to save content to database: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	// send audio to frontend
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":   true,
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
	req.Header.Set("x-content-length", fmt.Sprintf("%d", len(audioFile)))

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

func getAudioGenerations(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	email := r.Context().Value(ClaimsKey).(jwt.MapClaims)["email"].(string)

	audioGenerationRepo := model.AudioGenerationRepository{}
	audioGenerations, err := audioGenerationRepo.GetGenAudioHistory(email)
	if err != nil {
		http.Error(w, "Failed to get audio generations: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	json.NewEncoder(w).Encode(audioGenerations)
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
	http.Handle("/api/audio-generations", AuthMiddleware(http.HandlerFunc(getAudioGenerations)))

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", CorsMiddleware(http.DefaultServeMux))
}
