package repository

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"gitlab.com/back1ng1/question-bot/internal/database/entity"
)

type QuestionRepository struct {
	*pgx.Conn
}

func NewQuestionRepository(conn *pgx.Conn) *QuestionRepository {
	return &QuestionRepository{conn}
}

func (r QuestionRepository) FindQuestionsInPreset(presetId int) ([]entity.Question, error) {
	var questions []entity.Question

	rows, err := r.Query(
		context.Background(),
		`SELECT id, preset_id, title 
		FROM questions 
		WHERE preset_id=@preset_id`,
		pgx.NamedArgs{
			"preset_id": presetId,
		},
	)

	if err != nil {
		log.Fatal(err)
		return questions, err
	}

	defer rows.Close()

	for rows.Next() {
		question := entity.Question{}
		rows.Scan(&question.ID, &question.PresetId, &question.Title)
		questions = append(questions, question)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return questions, err
	}

	return questions, nil
}

func (r QuestionRepository) StoreQuestion(q entity.Question) error {
	row := r.QueryRow(
		context.Background(),
		`INSERT INTO questions(preset_id, title) 
		VALUES(@preset_id, @title) 
		RETURNING id, preset_id, title`,
		pgx.NamedArgs{
			"preset_id": q.PresetId,
			"title":     q.Title,
		},
	)

	var question entity.Question
	err := row.Scan(&question.ID, &question.PresetId, &question.Title)

	if err != nil {
		return err
	}

	return nil
}

func (r QuestionRepository) UpdateQuestionTitle(id int, q entity.Question) error {
	if len(q.Title) == 0 {
		return errors.New("title is null in update question")
	}

	commandTag, err := r.Exec(
		context.Background(),
		`UPDATE questions SET title=@title WHERE id=@id`,
		pgx.NamedArgs{
			"title": q.Title,
			"id":    id,
		},
	)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("cannot update question")
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
		return err
	}

	if commandTag.RowsAffected() != 1 {
		return errors.New("cannot delete question")
	}

	return nil
}
