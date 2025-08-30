package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
