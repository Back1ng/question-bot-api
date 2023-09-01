package crud_question

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
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
		logger.Log.Errorf(
			"app.usecase.crud_question.implement.GetByPreset() - uc.questionRepo.GetByPreset(presetId): %v. presetId: %d",
			err,
			presetId,
		)

		return nil, err
	}

	return out, err
}

func (uc *usecase) Create(in entity.Question) (*entity.Question, error) {
	out, err := uc.questionRepo.Create(in)

	if err != nil {
		logger.Log.Errorf(
			"app.usecase.crud_question.implement.Create() - uc.questionRepo.Create(in): %v. in: %#+v",
			err,
			in,
		)

		return nil, err
	}

	return out, nil
}

func (uc *usecase) Update(in entity.Question) (*entity.Question, error) {
	out, err := uc.questionRepo.Update(in)

	if err != nil {
		logger.Log.Errorf(
			"app.usecase.crud_question.implement.Update() - uc.questionRepo.Update(in): %v. in: %#+v",
			err,
			in,
		)

		return nil, err
	}

	return out, nil
}

func (uc *usecase) Delete(id int) error {
	err := uc.questionRepo.Delete(id)

	if err != nil {
		logger.Log.Errorf(
			"app.usecase.crud_question.implement.Delete() - uc.questionRepo.Delete(id): %v. id: %d",
			err,
			id,
		)
	}

	return err
}
