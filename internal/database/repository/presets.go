package repository

import (
	"context"

	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
	"gitlab.com/back1ng1/question-bot-api/pkg/postgres"
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
		logger.Log.Errorf("PresetRepository.FindPresets - r.Select: %v", err)
		return presets, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
	)

	if err != nil {
		logger.Log.Errorf("PresetRepository.FindPresets - r.Query: %v", err)
		return presets, err
	}

	defer rows.Close()

	for rows.Next() {
		var preset entity.Preset
		rows.Scan(&preset.ID, &preset.Title)
		presets = append(presets, preset)
	}

	if err := rows.Err(); err != nil {
		logger.Log.Errorf("PresetRepository.FindPresets - rows.Err: %v", err)
		return presets, err
	}

	return presets, nil
}

func (r PresetRepository) StorePreset(p entity.Preset) (entity.Preset, error) {
	sql, args, err := r.Insert("presets").
		Columns("title").
		Values(p.Title).
		Suffix("RETURNING id, title").
		ToSql()

	if err != nil {
		logger.Log.Errorf("PresetRepository.StorePreset - r.Insert: %v", err)
		return p, err
	}

	row := r.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	var preset entity.Preset
	if err := row.Scan(&preset.ID, &preset.Title); err != nil {
		logger.Log.Errorf("PresetRepository.StorePreset - rows.Scan: %v", err)
		return p, err
	}

	return preset, nil
}

func (r PresetRepository) UpdatePreset(id int, p entity.Preset) error {
	if len(p.Title) == 0 {
		logger.Log.Errorf("PresetRepository.UpdatePreset - len(p.title): %v", UpdatePresetsEmptyTitle)
		return UpdatePresetsEmptyTitle
	}

	sql, args, err := r.Update("presets").
		Set("title", p.Title).Where("id = ?", id).
		ToSql()

	if err != nil {
		logger.Log.Errorf("PresetRepository.UpdatePreset - r.Update: %v", err)
		return err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		logger.Log.Errorf("PresetRepository.UpdatePreset - r.Exec: %v", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf("PresetRepository.UpdatePreset - r.Update: %v", UpdatePresetsError)
		return UpdatePresetsError
	}

	return nil
}

func (r PresetRepository) DeletePreset(id int) error {
	sql, args, err := r.Delete("presets").
		Where("id = ?", id).
		ToSql()

	if err != nil {
		logger.Log.Errorf("PresetRepository.DeletePreset - r.Delete: %v", err)
		return err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		logger.Log.Errorf("PresetRepository.DeletePreset - r.Exec: %v", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf("PresetRepository.DeletePreset - r.Exec: %v", DeletePresetsError)
		return DeletePresetsError
	}

	return nil
}
