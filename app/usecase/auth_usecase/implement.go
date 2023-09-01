package auth_usecase

import (
	"errors"
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/pkg/tgauth"
)

type usecase struct {
	tokensRepo repository.TokenRepository
}

func NewUsecase(tokensRepo repository.TokenRepository) UseCase {
	return &usecase{
		tokensRepo: tokensRepo,
	}
}

func (uc *usecase) Authenticate(auth tgauth.Auth) (*string, error) {
	if auth.IsOutdated() {
		return nil, errors.New("token is outdated")
	}

	token, err := uc.tokensRepo.Get(auth.Hash)
	if err != nil {
		// token not find
		return nil, err
	}

	if token != nil {
		return &token.Hash, nil
	}

	if !auth.IsValid() {
		return nil, errors.New("auth usecase: auth not valid")
	}

	createdToken, err := uc.tokensRepo.Create(auth)

	if err := uc.tokensRepo.DeleteExcept(*createdToken); err != nil {
		return nil, err
	}

	return &createdToken.Hash, nil
}
