package question_repository_v1

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

func NewRepository(db *pgx.Conn, sb squirrel.StatementBuilderType) irepository.QuestionRepository {
	return &repository{
		db: db,
		sb: sb,
	}
}

func (r *repository) GetByPreset(presetId int) ([]*entity.Question, error) {
	var questions []*entity.Question

	sql, args, err := r.sb.Select("id", "preset_id", "title").
		From("questions").
		Where("preset_id = ?", presetId).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		question := entity.Question{}
		rows.Scan(&question.ID, &question.PresetId, &question.Title)
		questions = append(questions, &question)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func (r *repository) Create(in entity.Question) (*entity.Question, error) {
	var question entity.Question

	sql, args, err := r.sb.Insert("questions").
		Columns("preset_id", "title").
		Values(in.PresetId, in.Title).
		Suffix("RETURNING id, preset_id, title").
		ToSql()

	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	err = row.Scan(&question.ID, &question.PresetId, &question.Title)

	if err != nil {
		return nil, err
	}

	return &question, nil

}

func (r *repository) Update(in entity.Question) (*entity.Question, error) {
	if len(in.Title) == 0 {
		return nil, errors.New("title is empty")
	}

	sql, args, err := r.sb.Update("questions").
		Set("title", in.Title).
		Where("id = ?", in.ID).
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
		return nil, errors.New("update questions: rows not affected")
	}

	return &in, nil

}

func (r *repository) Delete(id int) error {
	sql, args, err := r.sb.Delete("questions").
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
		return errors.New("delete question: rows not affected")
	}

	return nil
}