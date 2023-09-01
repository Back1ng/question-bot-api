package crud_user

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
)

type usecase struct {
	userRepo repository.UserRepository
}

func NewUseCase(userRepo repository.UserRepository) UseCase {
	return &usecase{
		userRepo: userRepo,
	}
}

func (uc *usecase) GetByInterval(interval int) ([]int64, error) {
	users, err := uc.userRepo.GetByInterval(interval)

	if err != nil {
		logger.Log.Errorf(
			"app.usecase.crud_user.implement.GetByInterval() - uc.userRepo.GetByInterval(interval): %v. interval: %d",
			err,
			interval,
		)

		return nil, err
	}

	return users, nil
}

func (uc *usecase) GetByChatId(chatId int) (*entity.User, error) {
	user, err := uc.userRepo.GetByChatId(chatId)

	if err != nil {
		logger.Log.Errorf(
			"app.usecase.crud_user.implement.GetByChatId() - uc.userRepo.GetByChatId(chatId): %v. chatId: %d",
			err,
			chatId,
		)

		return nil, err
	}

	return user, nil
}

func (uc *usecase) Create(in entity.User) (*entity.User, error) {
	user, err := uc.userRepo.Create(in)

	if err != nil {
		logger.Log.Errorf(
			"app.usecase.crud_user.implement.Create() - uc.userRepo.Create(in): %v. in: %#+v",
			err,
			in,
		)

		return nil, err
	}

	return user, nil
}

func (uc *usecase) Update(in entity.User) (*entity.User, error) {
	user, err := uc.userRepo.Update(in)

	if err != nil {
		logger.Log.Errorf(
			"app.usecase.crud_user.implement.Update() - uc.userRepo.Update(in): %v. in: %#+v",
			err,
			in,
		)

		return nil, err
	}

	return user, nil
}
