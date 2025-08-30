package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"echogen/backend/config"
)

func GenAudio(finalContent string, req ttsRequest) (*http.Response, error) {
	OutputFormat := "mp3_44100_128"

	if req.VoiceId == "" {
		req.VoiceId = "3gsg3cxXyFLcGIfNbM6C"
	}
	if req.ModelId == "" {
		req.ModelId = "eleven_v3"
	}

	body, _ := json.Marshal(map[string]any{
		"text":     finalContent,
		"model_id": req.ModelId,
	})
	elevenURL := "https://api.elevenlabs.io/v1/text-to-speech/" + req.VoiceId + "?output_format=" + OutputFormat
	reqEleven, _ := http.NewRequest(http.MethodPost, elevenURL, bytes.NewReader(body))
	reqEleven.Header.Set("xi-api-key", config.ElevenAPIKey)
	reqEleven.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(reqEleven)
	if err != nil {
		return nil, fmt.Errorf("error calling elevenlabs: %w", err)
	}

	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("elevenlabs: %s", string(b))
	}

	return resp, nil
}
