package crud_question

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type usecase struct {
	question_repo repository.QuestionRepository
}

func NewUseCase(question_repo repository.QuestionRepository) UseCase {
	return &usecase{question_repo: question_repo}
}

func (uc *usecase) GetByPreset(presetId int) ([]*entity.Question, error) {
	out, err := uc.question_repo.GetByPreset(presetId)

	if err != nil {
		return nil, err
	}

	return out, err
}

func (uc *usecase) Create(in entity.Question) (*entity.Question, error) {
	out, err := uc.question_repo.Create(in)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Update(in entity.Question) (*entity.Question, error) {
	out, err := uc.question_repo.Update(in)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Delete(id int) error {
	return uc.question_repo.Delete(id)
}
