package crud_question

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type usecase struct {
	questionRepo repository.QuestionRepository
}

func NewUseCase(questionRepo repository.QuestionRepository) UseCase {
	return &usecase{questionRepo: questionRepo}
}

func (uc *usecase) GetByPreset(presetId int) ([]*entity.Question, error) {
	out, err := uc.questionRepo.GetByPreset(presetId)

	if err != nil {
		return nil, err
	}

	return out, err
}

func (uc *usecase) Create(in entity.Question) (*entity.Question, error) {
	out, err := uc.questionRepo.Create(in)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Update(in entity.Question) (*entity.Question, error) {
	out, err := uc.questionRepo.Update(in)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Delete(id int) error {
	return uc.questionRepo.Delete(id)
}
