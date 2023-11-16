package preset_repository_v1

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/Masterminds/squirrel"
	irepository "gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
)

type repository struct {
	db *pgxpool.Pool
	sb squirrel.StatementBuilderType
}

func NewRepository(db *pgxpool.Pool, sb squirrel.StatementBuilderType) irepository.PresetRepository {
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
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.GetAll() - r.sb.Select(): %v",
			err,
		)
		return nil, err
	}

	rows, err := r.db.Query(
		context.Background(),
		sql,
	)
	if err != nil {
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.GetAll() - r.db.Query(): %v",
			err,
		)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var preset entity.Preset
		rows.Scan(&preset.ID, &preset.Title)
		presets = append(presets, &preset)
	}

	if err := rows.Err(); err != nil {
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.GetAll() - rows.Err(): %v",
			err,
		)
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
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.Create() - r.sb.Insert(): %v. Preset: %#+v",
			err,
			in,
		)
		return nil, err
	}

	row := r.db.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	var preset entity.Preset
	if err := row.Scan(&preset.ID, &preset.Title); err != nil {
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.Create() - row.Scan(): %v. Preset: %#+v",
			err,
			in,
		)
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
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.Update() - r.sb.Update(): %v. Preset: %#+v",
			err,
			in,
		)
		return nil, err
	}

	commandTag, err := r.db.Exec(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.Update() - r.db.Exec(): %v. Preset: %#+v",
			err,
			in,
		)
		return nil, err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.Update() - commandTag.RowsAffected(): %v. Preset: %#+v",
			"update preset: no afftected rows",
			in,
		)
		return nil, errors.New("update preset: no afftected rows")
	}

	return &in, nil
}

func (r *repository) Delete(id int64) error {
	sql, args, err := r.sb.Delete("presets").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.Delete() - r.sb.Delete(): %v. Id: %d",
			err,
			id,
		)
		return err
	}

	commandTag, err := r.db.Exec(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.Delete() - r.db.Exec(): %v. Id: %d",
			err,
			id,
		)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf(
			"repository.preset_repository_v1.repository.Delete() - r.sb.Delete(): %v. Id: %d",
			"delete preset: no affected rows",
			id,
		)
		return errors.New("delete preset: no affected rows")
	}

	return nil
}
