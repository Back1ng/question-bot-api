package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"gitlab.com/back1ng1/question-bot/internal/database/entity"
)

type PresetRepository struct {
	*pgx.Conn
}

func NewPresetRepository(conn *pgx.Conn) *PresetRepository {
	return &PresetRepository{conn}
}

func (r PresetRepository) FindPresets() ([]entity.Preset, error) {
	var presets []entity.Preset
	rows, err := r.Query(
		context.Background(),
		`SELECT id, title FROM presets`,
	)

	if err != nil {
		log.Fatal(err)
		return presets, err
	}

	defer rows.Close()

	for rows.Next() {
		var preset entity.Preset
		rows.Scan(&preset.ID, &preset.Title)
		presets = append(presets, preset)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return presets, err
	}

	return presets, nil
}

func (r PresetRepository) StorePreset(p entity.Preset) error {
	row := r.QueryRow(
		context.Background(),
		"INSERT INTO presets(title) VALUES(@title) RETURNING id, title",
		pgx.NamedArgs{
			"title": p.Title,
		},
	)

	var preset entity.Preset
	if err := row.Scan(&preset.ID, &preset.Title); err != nil {
		return err
	}

	return nil
}

func (r PresetRepository) UpdatePreset(id int, p entity.Preset) error {
	if len(p.Title) == 0 {
		return errors.New("title is null in update preset")
	}

	commandTag, err := r.Exec(
		context.Background(),
		`UPDATE presets SET title=@title WHERE id=@id`,
		pgx.NamedArgs{
			"title": p.Title,
			"id":    id,
		},
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("cannot update presets")
	}

	return nil
}

func (r PresetRepository) DeletePreset(id int) error {
	commandTag, err := r.Exec(
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
		return errors.New("cannot delete preset")
	}

	return nil
}
