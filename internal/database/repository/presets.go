package repository

import (
	"context"
	"errors"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/postgres"
	"log"

	"github.com/jackc/pgx/v5"
)

type PresetRepository struct {
	postgres.PgConfig
}

func NewPresetRepository(pg postgres.PgConfig) *PresetRepository {
	return &PresetRepository{pg}
}

func (r PresetRepository) FindPresets() ([]entity.Preset, error) {
	var presets []entity.Preset

	sql, _, err := r.Select("id", "title").
		From("presets").
		ToSql()

	if err != nil {
		return presets, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
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
	sql, args, err := r.Insert("presets").
		Columns("title").
		Values(p.Title).
		Suffix("RETURNING id, title").
		ToSql()

	if err != nil {
		return err
	}

	row := r.QueryRow(
		context.Background(),
		sql,
		args...,
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
