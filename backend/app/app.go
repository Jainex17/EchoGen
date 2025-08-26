package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ttsRequest struct {
	Text         string `json:"text"`
	VoiceId      string `json:"voice_id"`
	ModelId      string `json:"model_id"`
	OutputFormat string `json:"output_format"`
}

var apiKey string

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

	if req.VoiceId == "" || req.ModelId == "" {
		fmt.Println("voice_id: ", req.VoiceId)
		fmt.Println("model_id: ", req.ModelId)
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	if req.OutputFormat == "" {
		req.OutputFormat = "mp3_44100_128"
	}

	body, _ := json.Marshal(map[string]any{
		"text":     req.Text,
		"model_id": req.ModelId,
	})

	elevenURL := "https://api.elevenlabs.io/v1/text-to-speech/" + req.VoiceId + "?output_format=" + req.OutputFormat
	reqEleven, _ := http.NewRequest(http.MethodPost, elevenURL, bytes.NewReader(body))
	reqEleven.Header.Set("xi-api-key", apiKey)
	reqEleven.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(reqEleven)
	if err != nil {
		http.Error(w, "elevenlabs error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		http.Error(w, "elevenlabs: "+string(b), resp.StatusCode)
		return
	}

	// store audio to file
	audio, err := os.Create("audio.mp3")
	if err != nil {
		http.Error(w, "error creating audio file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer audio.Close()

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Cache-Control", "no-store")

	// copy response body to both file and http response
	mw := io.MultiWriter(audio, w)
	if _, err := io.Copy(mw, resp.Body); err != nil {
		http.Error(w, "error copying audio: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey = os.Getenv("ELEVEN_API_KEY")
	if apiKey == "" {
		log.Fatal("ELEVEN_API_KEY is not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/api/tts", tts)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
