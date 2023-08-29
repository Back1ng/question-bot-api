package preset_repository_v1

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	irepository "gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type repository struct {
	db *pgx.Conn
	sb squirrel.StatementBuilderType
}

func NewRepository(db *pgx.Conn, sb squirrel.StatementBuilderType) irepository.PresetRepository {
	return &repository{
		db: db,
		sb: sb,
	}
}

func (r *repository) GetAll() ([]*entity.Preset, error) {
	var presets []*entity.Preset

	sql, _, err := r.sb.Select("id", "title").
		From("presets").
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(
		context.Background(),
		sql,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var preset entity.Preset
		rows.Scan(&preset.ID, &preset.Title)
		presets = append(presets, &preset)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return presets, nil
}

func (r *repository) Create(in entity.Preset) (*entity.Preset, error) {
	sql, args, err := r.sb.Insert("presets").
		Columns("title").
		Values(in.Title).
		Suffix("RETURNING id, title").
		ToSql()

	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	var preset entity.Preset
	if err := row.Scan(&preset.ID, &preset.Title); err != nil {
		return nil, err
	}

	return &preset, nil
}

func (r *repository) Update(in entity.Preset) (*entity.Preset, error) {
	if len(in.Title) == 0 {
		return nil, errors.New("update preset: title is empty")
	}

	sql, args, err := r.sb.Update("presets").
		Set("title", in.Title).Where("id = ?", in.ID).
		ToSql()

	if err != nil {
		return nil, err
	}

	commandTag, err := r.db.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		return nil, err
	}

	if commandTag.RowsAffected() != 1 {
		return nil, errors.New("update preset: no afftected rows")
	}

	return &in, nil
}

func (r *repository) Delete(id int64) error {
	sql, args, err := r.sb.Delete("presets").
		Where("id = ?", id).
		ToSql()

	if err != nil {
		return err
	}

	commandTag, err := r.db.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("delete preset: no affected rows")
	}

	return nil
}
