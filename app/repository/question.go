package repository

import "gitlab.com/back1ng1/question-bot-api/entity"

type QuestionRepository interface {
	GetByPreset(presetId int) ([]*entity.Question, error)
	Create(in entity.Question) (*entity.Question, error)
	Update(in entity.Question) (*entity.Question, error)
	Delete(id int) error
}
