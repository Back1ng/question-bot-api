package repository

import (
	"context"
	"errors"
	"gitlab.com/back1ng1/question-bot-api/internal/services/tgauth"
	"gitlab.com/back1ng1/question-bot-api/pkg/postgres"
)

type AuthRepository struct {
	postgres.PgConfig
}

func NewAuthRepository(pg postgres.PgConfig) *AuthRepository {
	return &AuthRepository{pg}
}

func (r AuthRepository) HasToken(hash string) (bool, error) {
	sql, args, err := r.Select("auth_date", "hash").
		From("tokens").
		Where("hash = ?", hash).
		ToSql()
	if err != nil {
		return false, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
		args...,
	)
	defer rows.Close()
	if err != nil {
		return false, err
	}

	var tokens []tgauth.Auth
	for rows.Next() {
		var auth tgauth.Auth
		if err := rows.Scan(&auth.AuthDate, &auth.Hash); err != nil {
			return false, err
		}
		tokens = append(tokens, auth)
	}

	for _, v := range tokens {
		if v.IsOutdated() {
			continue
		}

		return v.Hash != "", nil
	}

	return false, nil
}

func (r AuthRepository) GenerateToken(auth tgauth.Auth) (string, error) {
	if auth.IsOutdated() {
		return "", errors.New("auth: data is outdated")
	}

	hasToken, err := r.HasToken(auth.Hash)
	if err != nil {
		return "", err
	}

	if hasToken {
		return auth.Hash, nil
	}

	if auth.IsValid() {
		sql, args, err := r.Insert("tokens").
			Columns("auth_date", "first_name", "hash", "user_id", "username").
			Values(auth.AuthDate, auth.FirstName, auth.Hash, auth.Id, auth.Username).ToSql()
		if err != nil {
			return "", err
		}

		commandTag, err := r.Exec(
			context.Background(),
			sql,
			args...,
		)
		if err != nil {
			return "", err
		}

		if commandTag.RowsAffected() != 1 {
			return "", errors.New("auth: cannot store token")
		}

		return auth.Hash, nil
	}

	return "", errors.New("failed check auth")
}
