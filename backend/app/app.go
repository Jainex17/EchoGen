package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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
		} `json:"content"`
	} `json:"candidates"`
}

var apiKey string
var geminiAPIKey string

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

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Cache-Control", "no-store")

	mw := io.Writer(w)
	if _, err := io.Copy(mw, resp.Body); err != nil {
		http.Error(w, "error copying audio: "+err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("Audio saved to audio.mp3")
}

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey = os.Getenv("ELEVEN_API_KEY")
	geminiAPIKey = os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("ELEVEN_API_KEY is not set")
	}
	if geminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY is not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})

	http.HandleFunc("/api/tts", tts)
	http.HandleFunc("/auth/google/login", handleGoogleLogin)
	http.HandleFunc("/auth/google/callback", handleGoogleCallback)
	http.HandleFunc("/auth/me", handleProfile)
	http.HandleFunc("/auth/logout", handleLogout)

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", CorsMiddleware(http.DefaultServeMux))
}
