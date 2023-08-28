package answer_repository_v1

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

func New(db *pgx.Conn, sb squirrel.StatementBuilderType) irepository.AnswerRepository {
	return &repository{db: db, sb: sb}
}

func (r *repository) Get(questionId int) ([]*entity.Answer, error) {
	var answers []*entity.Answer

	sql, args, err := r.sb.Select("*").
		From("answers").
		Where("question_id = ?", questionId).
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
		var answer *entity.Answer
		rows.Scan(&answer.ID, &answer.QuestionId, &answer.Title, &answer.IsCorrect)
		answers = append(answers, answer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return answers, nil
}

func (r *repository) Create(in entity.Answer) (*entity.Answer, error) {
	sql, args, err := r.sb.Insert("answers").
		Columns("question_id", "title", "is_correct").
		Values(in.QuestionId, in.Title, in.IsCorrect).
		Suffix("RETURNING id, is_correct").
		ToSql()

	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	if err := row.Scan(&in.ID, &in.IsCorrect); err != nil {
		return nil, err
	}

	return &in, nil

}

func (r *repository) Update(in entity.Answer) (*entity.Answer, error) {
	if len(in.Title) == 0 {
		return nil, errors.New("empty answer title")
	}

	sql, args, err := r.sb.Update("answers").
		Set("title", in.Title).
		Set("is_correct", in.IsCorrect).
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
		return nil, errors.New("cannot update answer")
	}

	return &in, nil
}

func (r *repository) Delete(id int) error {
	if id == 0 {
		return errors.New("id answer not presented")
	}

	sql, args, err := r.sb.Delete("answers").
		Where("id = ?", id).ToSql()

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
		return errors.New("cannot delete answer")
	}

	return nil
}
