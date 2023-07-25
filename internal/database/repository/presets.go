package repository

import (
	"context"

	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
)

func StorePreset(p models.Preset) (models.Preset, error) {
	row := database.Database.DB.QueryRow(
		context.Background(),
		"INSERT INTO presets(title) VALUES($1) RETURNING id, title",
		p.Title,
	)

	preset := models.Preset{}
	err := row.Scan(&preset.ID, &preset.Title)

	if err != nil {
		return preset, err
	}

	return preset, nil
}
