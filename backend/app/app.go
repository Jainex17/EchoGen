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

var SystemPrompt = `You are a podcast narrator. The user will provide a topic, and you must create a podcast-style script about it.  

Rules:  
- Keep the script concise: about 30–60 seconds of spoken audio (roughly 120–150 words).  
- Write only the podcast content — do not include meta-text like “I understand” or “Here’s your explanation.”  
- Use a natural, engaging, conversational tone as if youre speaking to an audience.  
- If the user provides a podcast/host name, use it naturally in the introduction.  
- If no name is provided, make up a short, catchy podcast name and use it.  
- Structure:  
  - Intro hook (grab attention).  
  - Main explanation (clear, flowing, engaging).  
  - Closing note (wrap up neatly).  
- Avoid robotic or overly formal wording — keep it human and story-like.  
`

var apiKey string
var geminiAPIKey string

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func tts(w http.ResponseWriter, r *http.Request) {
	enableCors(w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

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

	if req.Prompt == "" {
		http.Error(w, "Missing prompt", http.StatusBadRequest)
		return
	}

	if req.VoiceId == "" {
		req.VoiceId = "3gsg3cxXyFLcGIfNbM6C"
	}
	if req.ModelId == "" {
		req.ModelId = "eleven_v3"
	}

	OutputFormat := "mp3_44100_128"

	// prompt to content
	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{
						"text": SystemPrompt + "\n\nuser: " + req.Prompt,
					},
				},
			},
		},
	}
	contentBody, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "error marshaling payload: "+err.Error(), http.StatusInternalServerError)
		return
	}

	geminiURL := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"
	reqGemini, err := http.NewRequest(http.MethodPost, geminiURL+"?key="+geminiAPIKey, bytes.NewReader(contentBody))
	if err != nil {
		log.Fatal("error creating request:", err)
	}
	reqGemini.Header.Set("Content-Type", "application/json")

	respGemini, err := http.DefaultClient.Do(reqGemini)
	if err != nil {
		http.Error(w, "gemini error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer respGemini.Body.Close()

	respBody, _ := io.ReadAll(respGemini.Body)
	var gemResp GeminiResponse
	err = json.Unmarshal(respBody, &gemResp)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	finalContent := ""
	if len(gemResp.Candidates) > 0 && len(gemResp.Candidates[0].Content.Parts) > 0 {
		finalContent = gemResp.Candidates[0].Content.Parts[0].Text
	}

	fmt.Println("Final Content:\n", finalContent)

	// text to speech
	body, _ := json.Marshal(map[string]any{
		"text":     finalContent,
		"model_id": req.ModelId,
	})
	elevenURL := "https://api.elevenlabs.io/v1/text-to-speech/" + req.VoiceId + "?output_format=" + OutputFormat
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
	audio, err := os.Create("audio.mp3_" + uuid.New().String())
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

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
