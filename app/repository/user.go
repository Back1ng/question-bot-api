package repository

import "gitlab.com/back1ng1/question-bot-api/entity"

type UserRepository interface {
	GetByInterval(interval int) ([]int64, error)
	GetByChatId(chatId int) (*entity.User, error)
	Create(in entity.User) (*entity.User, error)
	Update(in entity.User) (*entity.User, error)
}
