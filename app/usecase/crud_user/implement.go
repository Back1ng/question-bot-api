package crud_user

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
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
		return nil, err
	}

	return users, nil
}

func (uc *usecase) GetByChatId(chatId int) (*entity.User, error) {
	user, err := uc.userRepo.GetByChatId(chatId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *usecase) Create(in entity.User) (*entity.User, error) {
	user, err := uc.userRepo.Create(in)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *usecase) Update(in entity.User) (*entity.User, error) {
	user, err := uc.userRepo.Update(in)

	if err != nil {
		return nil, err
	}

	return user, nil
}
