package model

import (
	"time"

	"echogen/backend/config"
)

type AudioGeneration struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Prompt    string    `db:"prompt"`
	Content   string    `db:"content"`
	AudioURL  string    `db:"audio_url"`
	CreatedAt time.Time `db:"created_at"`
}

type AudioGenerationRepository struct{}

func (r *AudioGenerationRepository) CreateAudioGeneration(audioGeneration *AudioGeneration) error {
	query := `
		INSERT INTO audio_generations (user_id, prompt, content, audio_url) VALUES ($1, $2, $3, $4)
	`

	_, err := config.DB.Exec(query, audioGeneration.UserID, audioGeneration.Prompt, audioGeneration.Content, audioGeneration.AudioURL)
	return err
}
