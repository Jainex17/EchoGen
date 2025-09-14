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

type audioHistory struct {
	Prompt   string `db:"prompt"`
	AudioURL string `db:"audio_url"`
}

type AudioGenerationRepository struct{}

func (r *AudioGenerationRepository) CreateAudioGeneration(audioGeneration *AudioGeneration) error {
	query := `
		INSERT INTO audio_generations (user_id, prompt, content, audio_url) VALUES ($1, $2, $3, $4)
	`

	_, err := config.DB.Exec(query, audioGeneration.UserID, audioGeneration.Prompt, audioGeneration.Content, audioGeneration.AudioURL)
	return err
}

func (r *AudioGenerationRepository) GetGenAudioHistory(email string) ([]audioHistory, error) {

	userRepo := UserRepository{}
	resdata := []audioHistory{}
	var data AudioGeneration

	user, err := userRepo.GetUserByEmail(email)
	if err != nil {
		return resdata, err
	}

	query := `
		SELECT prompt, audio_url FROM audio_generations WHERE 
		user_id = $1
		ORDER BY created_at DESC
		LIMIT 10
	`

	rows, err := config.DB.Query(query, user.ID)
	if err != nil {
		return resdata, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&data.Prompt, &data.AudioURL)
		if err != nil {
			return resdata, err
		}
		resdata = append(resdata, audioHistory{
			Prompt:   data.Prompt,
			AudioURL: data.AudioURL,
		})
	}

	return resdata, nil
}
