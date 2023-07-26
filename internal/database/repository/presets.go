package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
)

func FindPresets() ([]models.Preset, error) {
	presets := []models.Preset{}
	rows, err := database.Database.DB.Query(
		context.Background(),
		`SELECT id, title FROM presets`,
	)

	if err != nil {
		log.Fatal(err)
		return presets, err
	}

	defer rows.Close()

	for rows.Next() {
		preset := models.Preset{}
		rows.Scan(&preset.ID, &preset.Title)
		presets = append(presets, preset)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return presets, err
	}

	return presets, nil
}

func StorePreset(p models.Preset) (models.Preset, error) {
	row := database.Database.DB.QueryRow(
		context.Background(),
		"INSERT INTO presets(title) VALUES(@title) RETURNING id, title",
		pgx.NamedArgs{
			"title": p.Title,
		},
	)

	preset := models.Preset{}
	err := row.Scan(&preset.ID, &preset.Title)

	if err != nil {
		return preset, err
	}

	return preset, nil
}

func UpdatePreset(id int, p models.Preset) (models.Preset, error) {
	if len(p.Title) == 0 {
		return p, errors.New("title is null in update preset")
	}

	commandTag, err := database.Database.DB.Exec(
		context.Background(),
		`UPDATE presets SET title=@title WHERE id=@id`,
		pgx.NamedArgs{
			"title": p.Title,
			"id":    id,
		},
	)

	if err != nil {
		return p, err
	}

	if commandTag.RowsAffected() != 1 {
		return p, errors.New("cannot update presets")
	}

	return p, nil
}

func DeletePreset(id int) error {
	commandTag, err := database.Database.DB.Exec(
		context.Background(),
		`DELETE FROM presets WHERE id=@id`,
		pgx.NamedArgs{
			"id": id,
		},
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("cannot delete presets")
	}

	return nil
}
