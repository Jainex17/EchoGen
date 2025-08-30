package model

import "time"

type AudioGeneration struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Prompt    string    `db:"prompt"`
	Content   string    `db:"content"`
	AudioURL  string    `db:"audio_url"`
	CreatedAt time.Time `db:"created_at"`
}

type AudioGenerationRepository struct{}
