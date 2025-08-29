package app

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func GenAudio(finalContent string, req ttsRequest) (*http.Response, error) {
	OutputFormat := "mp3_44100_128"

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
		return nil, fmt.Errorf("error calling elevenlabs: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("elevenlabs: %s", string(b))
	}

	// store audio to file
	audio, err := os.Create("audio.mp3_" + uuid.New().String())
	if err != nil {
		return nil, fmt.Errorf("error creating audio file: %w", err)
	}
	defer audio.Close()

	mw := io.Writer(audio)
	if _, err := io.Copy(mw, resp.Body); err != nil {
		return nil, fmt.Errorf("error copying audio: %w", err)
	}

	return resp, nil
}
