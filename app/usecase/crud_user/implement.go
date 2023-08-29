package crud_user

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type usecase struct {
	user_repo repository.UserRepository
}

func NewUseCase(user_repo repository.UserRepository) UseCase {
	return &usecase{
		user_repo: user_repo,
	}
}

func (uc *usecase) GetByInterval(interval int) ([]int64, error) {
	users, err := uc.user_repo.GetByInterval(interval)

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *usecase) GetByChatId(chatId int) (*entity.User, error) {
	user, err := uc.user_repo.GetByChatId(chatId)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *usecase) Create(in entity.User) (*entity.User, error) {
	user, err := uc.user_repo.Create(in)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *usecase) Update(in entity.User) (*entity.User, error) {
	user, err := uc.user_repo.Update(in)

	if err != nil {
		return nil, err
	}

	return user, nil
}
