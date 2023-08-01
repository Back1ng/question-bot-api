package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"gitlab.com/back1ng1/question-bot/internal/database/entity"
)

type AnswerRepository struct {
	*pgx.Conn
}

func NewAnswerRepository(conn *pgx.Conn) *AnswerRepository {
	return &AnswerRepository{conn}
}

func (r *AnswerRepository) FindAnswersInQuestion(questionId int) ([]entity.Answer, error) {
	var answers []entity.Answer

	rows, err := r.Query(
		context.Background(),
		`SELECT * FROM answers WHERE question_id = @question_id`,
		pgx.NamedArgs{
			"question_id": questionId,
		},
	)
	if err != nil {
		return answers, err
	}

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
	row := r.QueryRow(
		context.Background(),
		`INSERT INTO answers(question_id, title, is_correct)
			VALUES(@question_id, @title, @is_correct)
			RETURNING id, is_correct`,
		pgx.NamedArgs{
			"question_id": answer.QuestionId,
			"title":       answer.Title,
			"is_correct":  answer.IsCorrect,
		},
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

	commandTag, err := r.Exec(
		context.Background(),
		`UPDATE answers 
			SET title=@title, is_correct=@is_correct 
			WHERE id=@id`,
		pgx.NamedArgs{
			"title":      answer.Title,
			"is_correct": answer.IsCorrect,
			"id":         answer.ID,
		},
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

	commandTag, err := r.Exec(
		context.Background(),
		`DELETE FROM answers WHERE id=@id`,
		pgx.NamedArgs{
			"id": answer.ID,
		},
	)

	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("cannot delete answer")
	}

	return nil
}
