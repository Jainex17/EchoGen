package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"echogen/backend/config"
)

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

func GenContent(req *ttsRequest) (string, error) {

	if req.Prompt == "" {
		return "", fmt.Errorf("missing prompt")
	}

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
		return "", err
	}

	geminiURL := "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"
	reqGemini, err := http.NewRequest(http.MethodPost, geminiURL+"?key="+config.GeminiAPIKey, bytes.NewReader(contentBody))
	if err != nil {
		log.Fatal("error creating request:", err)
	}
	reqGemini.Header.Set("Content-Type", "application/json")

	respGemini, err := http.DefaultClient.Do(reqGemini)
	if err != nil {
		return "", err
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

	return finalContent, nil
}
