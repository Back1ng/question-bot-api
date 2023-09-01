package tokens_repository_v1

import (
	"context"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	irepository "gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/pkg/tgauth"
)

type repository struct {
	db *pgx.Conn
	sb squirrel.StatementBuilderType
}

func NewRepository(db *pgx.Conn, sb squirrel.StatementBuilderType) irepository.TokenRepository {
	return &repository{
		db: db,
		sb: sb,
	}
}

func (r *repository) Get(hash string) (*tgauth.Auth, error) {
	sql, args, err := r.sb.Select("auth_date", "hash").
		From("tokens").
		Where("hash = ?", hash).
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
		var auth tgauth.Auth

		if err := rows.Scan(&auth.AuthDate, &auth.Hash); err != nil {
			return nil, err
		}

		return &auth, nil
	}

	return nil, nil
}

func (r *repository) Create(auth tgauth.Auth) (*tgauth.Auth, error) {
	sql, args, err := r.sb.Insert("tokens").
		Columns("auth_date", "first_name", "hash", "user_id", "username").
		Values(auth.AuthDate, auth.FirstName, auth.Hash, auth.Id, auth.Username).ToSql()

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
		return nil, errors.New("no rows affected at authRepository.Create()")
	}

	return &auth, nil
}

func (r *repository) DeleteExcept(auth tgauth.Auth) error {
	sql, args, err := r.sb.Delete("tokens").
		Where("hash != ?", auth.Hash).
		Where("user_id = ?", auth.Id).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		return err
	}

	return nil
}