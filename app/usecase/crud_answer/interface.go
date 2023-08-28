package crud_answers

import "gitlab.com/back1ng1/question-bot-api/entity"

type UseCase interface {
	Get(questionId int) ([]entity.Answer, error)
	Create(in entity.Answer) (entity.Answer, error)
	Update(in entity.Answer) (entity.Answer, error)
	Delete(id int) error
}
