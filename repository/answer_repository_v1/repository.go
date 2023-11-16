package answer_repository_v1

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

func NewRepository(db *pgxpool.Pool, sb squirrel.StatementBuilderType) irepository.AnswerRepository {
	return &repository{
		db: db,
		sb: sb,
	}
}

func (r *repository) Get(questionId int) ([]*entity.Answer, error) {
	var answers []*entity.Answer

	sql, args, err := r.sb.Select("*").
		From("answers").
		Where("question_id = ?", questionId).
		ToSql()
	if err != nil {
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Get() - r.sb.Select(\"*\"): %v",
			err,
		)

		return nil, err
	}

	rows, err := r.db.Query(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Get() - r.db.Query(): %v. sql: %#+v. args: %#+v",
			err,
			sql,
			args,
		)

		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var answer entity.Answer
		rows.Scan(&answer.ID, &answer.QuestionId, &answer.Title, &answer.IsCorrect)
		answers = append(answers, &answer)
	}

	if err := rows.Err(); err != nil {
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Get() - rows.Err(): %v",
			err,
		)

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
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Create() - r.sb.Insert(\"answers\"): %v",
			err,
		)

		return nil, err
	}

	row := r.db.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	if err := row.Scan(&in.ID, &in.IsCorrect); err != nil {
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Create() - row.Scan(): %v, sql: %#+v, args: %#+v",
			err,
			sql,
			args,
		)

		return nil, err
	}

	return &in, nil
}

func (r *repository) Update(in entity.Answer) (*entity.Answer, error) {
	if len(in.Title) == 0 {
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Update() - len(in.Title) == 0: %v. in: %#+v",
			"empty answer title",
			in,
		)

		return nil, errors.New("empty answer title")
	}

	sql, args, err := r.sb.Update("answers").
		Set("title", in.Title).
		Set("is_correct", in.IsCorrect).
		Where("id = ?", in.ID).
		ToSql()
	if err != nil {
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Update() - r.sb.Update(\"answers\"): %v",
			err,
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
			"repository.answer_repository_v1.repository.Update() - r.db.Exec(): %v. sql: %#+v. args: %#+v",
			err,
			sql,
			args,
		)

		return nil, err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Update() - commandTag.RowsAffected() != 1: %v, sql: %#+v, args: %#+v",
			"not affected rows",
			sql,
			args,
		)

		return nil, errors.New("cannot update answer")
	}

	return &in, nil
}

func (r *repository) Delete(id int) error {
	if id == 0 {
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Delete() - id == 0: %v",
			"id answer not presented",
		)

		return errors.New("id answer not presented")
	}

	sql, args, err := r.sb.Delete("answers").
		Where("id = ?", id).ToSql()
	if err != nil {
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Delete() - r.sb.Delete(\"answers\"): %v",
			err,
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
			"repository.answer_repository_v1.repository.Delete() - r.db.Exec(): %v. sql: %#+v, args %#+v",
			err,
			sql,
			args,
		)

		return err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf(
			"repository.answer_repository_v1.repository.Delete() - commandTag.RowsAffected() != 1: %v, sql: %#+v, args: %#+v",
			"not affected rows",
			sql,
			args,
		)

		return errors.New("cannot delete answer")
	}

	return nil
}
