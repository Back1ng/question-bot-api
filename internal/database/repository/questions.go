package repository

import (
	"context"

	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
	"gitlab.com/back1ng1/question-bot-api/pkg/postgres"

	"github.com/jackc/pgx/v5"
)

type QuestionRepository struct {
	postgres.PgConfig
}

func NewQuestionRepository(pg postgres.PgConfig) *QuestionRepository {
	return &QuestionRepository{pg}
}

func (r QuestionRepository) FindQuestionsInPreset(presetId int) ([]entity.Question, error) {
	var questions []entity.Question

	sql, args, err := r.Select("id", "preset_id", "title").
		From("questions").
		Where("preset_id = ?", presetId).
		ToSql()

	if err != nil {
		logger.Log.Errorf("QuestionRepository.FindQuestionsInPreset - r.Select: %v", err)
		return questions, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		logger.Log.Errorf("QuestionRepository.FindQuestionsInPreset - r.Query: %v", err)
		return questions, err
	}

	defer rows.Close()

	for rows.Next() {
		question := entity.Question{}
		rows.Scan(&question.ID, &question.PresetId, &question.Title)
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		logger.Log.Errorf("QuestionRepository.FindQuestionsInPreset - rows.Err: %v", err)
		return questions, err
	}

	return questions, nil
}

func (r QuestionRepository) StoreQuestion(q entity.Question) error {
	sql, args, err := r.Insert("questions").
		Columns("preset_id", "title").
		Values(q.PresetId, q.Title).
		Suffix("RETURNING id, preset_id, title").
		ToSql()

	if err != nil {
		logger.Log.Errorf("QuestionRepository.StoreQuestion - r.Insert: %v", err)
		return err
	}

	row := r.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	var question entity.Question
	err = row.Scan(&question.ID, &question.PresetId, &question.Title)

	if err != nil {
		logger.Log.Errorf("QuestionRepository.StoreQuestion - rows.Scan: %v", err)
		return err
	}

	return nil
}

func (r QuestionRepository) UpdateQuestionTitle(id int, q entity.Question) error {
	if len(q.Title) == 0 {
		logger.Log.Errorf("QuestionRepository.UpdateQuestionTitle - len(q.Title): %v", UpdateQuestionEmptyTitle)
		return UpdateQuestionEmptyTitle
	}

	sql, args, err := r.Update("questions").
		Set("title", q.Title).
		Where("id = ?", id).
		ToSql()
	if err != nil {
		logger.Log.Errorf("QuestionRepository.UpdateQuestionTitle - r.Update: %v", err)
		return err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf("QuestionRepository.UpdateQuestionTitle - r.Exec: %v", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf("QuestionRepository.UpdateQuestionTitle - r.Exec: %v", UpdateQuestionError)
		return UpdateQuestionError
	}

	return nil
}

func (r QuestionRepository) DeleteQuestion(id int) error {
	commandTag, err := r.Exec(
		context.Background(),
		`DELETE FROM questions WHERE id=@id`,
		pgx.NamedArgs{
			"id": id,
		},
	)

	if err != nil {
		logger.Log.Errorf("QuestionRepository.DeleteQuestion - r.Exec: %v", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf("QuestionRepository.DeleteQuestion - r.Exec: %v", DeleteQuestionError)
		return DeleteQuestionError
	}

	return nil
}
