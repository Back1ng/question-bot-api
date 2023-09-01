package auth_usecase

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
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
		logger.Log.Errorf(
			"app.usecase.auth_usecase.implement.Authenticate() - auth.IsOutdated(): %v",
			tgauth.TokenIsOutdated,
		)

		return nil, tgauth.TokenIsOutdated
	}

	token, err := uc.tokensRepo.Get(auth.Hash)
	if err != nil {
		logger.Log.Errorf(
			"app.usecase.auth_usecase.implement.Authenticate() - uc.tokensRepo.Get(auth.Hash): %v",
			err,
		)

		return nil, err
	}

	if token != nil {
		return &token.Hash, nil
	}

	if !auth.IsValid() {
		logger.Log.Errorf(
			"app.usecase.auth_usecase.implement.Authenticate() - !auth.IsValid(): %v",
			tgauth.AuthNotValid,
		)
		return nil, tgauth.AuthNotValid
	}

	createdToken, err := uc.tokensRepo.Create(auth)
	if err != nil {
		logger.Log.Errorf(
			"app.usecase.auth_usecase.implement.Authenticate() - uc.tokensRepo.Create(auth): %v",
			err,
		)
		return nil, err
	}

	if err := uc.tokensRepo.DeleteExcept(*createdToken); err != nil {
		logger.Log.Errorf(
			"app.usecase.auth_usecase.implement.Authenticate() - uc.tokensRepo.DeleteExcept(*createdToken): %v",
			err,
		)

		return nil, err
	}

	return &createdToken.Hash, nil
}
