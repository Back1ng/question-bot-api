package repository

import (
	"context"

	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
	"gitlab.com/back1ng1/question-bot-api/pkg/postgres"
	"gitlab.com/back1ng1/question-bot-api/pkg/tgauth"
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
		logger.Log.Errorf("AuthRepository.HasToken - r.Select: %v", err)
		return false, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf("AuthRepository.HasToken - r.Query: %v", err)
		return false, err
	}
	defer rows.Close()

	var tokens []tgauth.Auth
	for rows.Next() {
		var auth tgauth.Auth

		if err := rows.Scan(&auth.AuthDate, &auth.Hash); err != nil {
			logger.Log.Errorf("AuthRepository.HasToken - rows.Scan: %v", err)
			return false, err
		}

		tokens = append(tokens, auth)
	}

	for _, v := range tokens {
		if v.IsOutdated() {
			continue
		}

		if v.Hash == hash {
			return true, nil
		}
	}

	return false, nil
}

func (r AuthRepository) GenerateToken(auth tgauth.Auth) (string, error) {
	if auth.IsOutdated() {
		logger.Log.Errorf("AuthRepository.GenerateToken - auth.IsOutdated: Given token is outdated. TgAuth: %#+v", auth)
		return "", AuthDataIsOutdated
	}

	hasToken, err := r.HasToken(auth.Hash)
	if err != nil {
		logger.Log.Errorf("AuthRepository.GenerateToken - r.HasToken: %v. TgAuth: %#+v", err, auth)
		return "", err
	}

	if hasToken {
		return auth.Hash, nil
	}

	if !auth.IsValid() {
		logger.Log.Errorf("AuthRepository.GenerateToken - auth.IsValid: %v. TgAuth: %#+v", AuthFailedCheck, auth)
		return "", AuthFailedCheck
	}

	sql, args, err := r.Insert("tokens").
		Columns("auth_date", "first_name", "hash", "user_id", "username").
		Values(auth.AuthDate, auth.FirstName, auth.Hash, auth.Id, auth.Username).ToSql()
	if err != nil {
		logger.Log.Errorf("AuthRepository.GenerateToken - r.Insert: %v. TgAuth: %#+v", err, auth)
		return "", err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf("AuthRepository.GenerateToken - r.Exec: %v. TgAuth: %#+v", err, auth)
		return "", err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf("AuthRepository.GenerateToken - r.Exec: %v. TgAuth: %#+v", AuthCannotStoreToken, auth)
		return "", AuthCannotStoreToken
	}

	sql, args, err = r.Delete("tokens").
		Where("hash != ?", auth.Hash).
		Where("user_id = ?", auth.Id).
		ToSql()

	if err != nil {
		logger.Log.Errorf("AuthRepository.GenerateToken - r.Delete: %v. TgAuth: %#+v", err, auth)
		return "", err
	}

	_, err = r.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		logger.Log.Errorf("AuthRepository.GenerateToken - r.Exec: %v. TgAuth: %#+v", err, auth)
		return "", err
	}

	return auth.Hash, nil
}
