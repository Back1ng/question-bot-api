package repository

import "gitlab.com/back1ng1/question-bot-api/pkg/tgauth"

type TokenRepository interface {
	Get(hash string) (*tgauth.Auth, error)
	Create(auth tgauth.Auth) (*tgauth.Auth, error)
	DeleteExcept(auth tgauth.Auth) error
}
