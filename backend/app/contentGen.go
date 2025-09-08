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

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		}
	}
}

var PodCastPrompt = `Generate a professional, insightful discussion on the given topic. Keep the tone formal, structured, and clear as if two experts are addressing an audience. Stay under 100 words and under 1 minute of spoken content. Focus on delivering key points with authority and clarity.`
var YouTuberPrompt = `Generate a casual, energetic explanation on the given topic. Use simple language, friendly tone, and engaging style, like a YouTuber speaking to their audience. Stay under 100 words and under 1 minute of spoken content. Keep it fun, relatable, and easy to follow.`
var TutorialPrompt = `Generate a step-by-step tutorial on the given topic. Use a clear and instructional tone, breaking the explanation into simple, numbered steps. Stay under 100 words and under 1 minute of spoken content. Focus on practical, easy-to-apply guidance.`

func GenContent(req *ttsRequest) (string, error) {

	if req.Prompt == "" {
		return "", fmt.Errorf("missing prompt")
	}

	var systemPrompt string
	switch req.Style {
	case "podcast":
		systemPrompt = PodCastPrompt
	case "youtuber":
		systemPrompt = YouTuberPrompt
	case "tutorial":
		systemPrompt = TutorialPrompt
	default:
		return "", fmt.Errorf("invalid style")
	}

	payload := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]string{
					{
						"text": systemPrompt + "\n\nuser: " + req.Prompt,
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
