package repository

import (
	"context"

	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
	"gitlab.com/back1ng1/question-bot-api/pkg/postgres"
)

type AnswerRepository struct {
	postgres.PgConfig
}

func NewAnswerRepository(pg postgres.PgConfig) *AnswerRepository {
	return &AnswerRepository{pg}
}

func (r *AnswerRepository) FindAnswersInQuestion(questionId int) ([]entity.Answer, error) {
	var answers []entity.Answer

	sql, args, err := r.Select("*").
		From("answers").
		Where("question_id = ?", questionId).
		ToSql()

	if err != nil {
		logger.Log.Errorf("AnswerRepository.FindAnswersInQuestion - r.Select: %v", err)
		return answers, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf("AnswerRepository.FindAnswersInQuestion - r.Query: %v", err)
		return answers, err
	}

	defer rows.Close()

	for rows.Next() {
		var answer entity.Answer
		rows.Scan(&answer.ID, &answer.QuestionId, &answer.Title, &answer.IsCorrect)
		answers = append(answers, answer)
	}

	if err := rows.Err(); err != nil {
		logger.Log.Errorf("AnswerRepository.FindAnswersInQuestion - rows.Err: %v", err)
		return answers, err
	}

	return answers, nil
}

func (r *AnswerRepository) StoreAnswer(answer entity.Answer) (entity.Answer, error) {
	sql, args, err := r.Insert("answers").
		Columns("question_id", "title", "is_correct").
		Values(answer.QuestionId, answer.Title, answer.IsCorrect).
		Suffix("RETURNING id, is_correct").
		ToSql()
	if err != nil {
		logger.Log.Errorf("AnswerRepository.StoreAnswer - r.Insert: %v", err)
		return answer, err
	}

	row := r.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	if err := row.Scan(&answer.ID, &answer.IsCorrect); err != nil {
		logger.Log.Errorf("AnswerRepository.StoreAnswer - row.Scan: %v", err)
		return answer, err
	}

	return answer, nil
}

func (r *AnswerRepository) UpdateAnswer(answer entity.Answer) error {
	if len(answer.Title) == 0 {
		logger.Log.Errorf("AnswerRepository.UpdateAnswer - answer.Title: %v", UpdateAnswerEmptyTitle)
		return UpdateAnswerEmptyTitle
	}

	sql, args, err := r.Update("answers").
		Set("title", answer.Title).
		Set("is_correct", answer.IsCorrect).
		Where("id = ?", answer.ID).
		ToSql()

	if err != nil {
		logger.Log.Errorf("AnswerRepository.UpdateAnswer - r.Update: %v", err)
		return err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		logger.Log.Errorf("AnswerRepository.UpdateAnswer - r.Exec: %v", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf("AnswerRepository.UpdateAnswer - r.Exec: %v", UpdateAnswerError)
		return UpdateAnswerError
	}

	return nil
}

func (r *AnswerRepository) DeleteAnswer(answer entity.Answer) error {
	if answer.ID == 0 {
		logger.Log.Errorf("AnswerRepository.DeleteAnswer - answer.ID: %v", DeleteAnswerIdNotPresented)
		return DeleteAnswerIdNotPresented
	}

	sql, args, err := r.Delete("answers").
		Where("id = ?", answer.ID).ToSql()

	if err != nil {
		logger.Log.Errorf("AnswerRepository.DeleteAnswer - r.Delete: %v", err)
		return err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		logger.Log.Errorf("AnswerRepository.DeleteAnswer - r.Exec: %v", err)
		return err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf("AnswerRepository.DeleteAnswer - r.Exec: %v", DeleteAnswerError)
		return DeleteAnswerError
	}

	return nil
}
