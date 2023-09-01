package auth_usecase

import "gitlab.com/back1ng1/question-bot-api/pkg/tgauth"

type UseCase interface {
	Authenticate(auth tgauth.Auth) (*string, error)
}
