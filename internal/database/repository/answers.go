package repository

import (
	"context"
	"errors"

	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
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
		return answers, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		return answers, err
	}

	defer rows.Close()

	for rows.Next() {
		var answer entity.Answer
		rows.Scan(&answer.ID, &answer.QuestionId, &answer.Title, &answer.IsCorrect)
		answers = append(answers, answer)
	}

	if err := rows.Err(); err != nil {
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
		return answer, err
	}

	row := r.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	if err := row.Scan(&answer.ID, &answer.IsCorrect); err != nil {
		return answer, err
	}

	return answer, nil
}

func (r *AnswerRepository) UpdateAnswer(answer entity.Answer) error {
	if len(answer.Title) == 0 {
		return errors.New("title is null in update answer")
	}

	sql, args, err := r.Update("answers").
		Set("title", answer.Title).
		Set("is_correct", answer.IsCorrect).
		Where("id", answer.ID).
		ToSql()

	if err != nil {
		return err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("cannot update answers")
	}

	return nil
}

func (r *AnswerRepository) DeleteAnswer(answer entity.Answer) error {
	if answer.ID == 0 {
		return errors.New("id is not presented")
	}

	sql, args, err := r.Delete("answers").Where("id = ?", answer.ID).ToSql()

	if err != nil {
		return err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("cannot delete answer")
	}

	return nil
}
