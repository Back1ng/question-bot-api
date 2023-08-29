package crud_answers

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
)

type usecase struct {
	answerRepo repository.AnswerRepository
}

func NewUseCase(answerRepo repository.AnswerRepository) UseCase {
	return &usecase{answerRepo: answerRepo}
}

func (uc *usecase) Get(questionId int) ([]*entity.Answer, error) {
	answers, err := uc.answerRepo.Get(questionId)

	if err != nil {
		logger.Log.Errorf("AnswerApi.GetAnswers - r.Repo.FindAnswersInQuestion: %v. QuestionId: %d", err, questionId)
		return nil, err
	}

	return answers, nil
}

func (uc *usecase) Create(in entity.Answer) (*entity.Answer, error) {
	out, err := uc.answerRepo.Create(in)

	if err != nil {
		logger.Log.Errorf("AnswerApi.StoreAnswers - r.Repo.StoreAnswer: %v. Answer: %#+v", err, in)
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Update(in entity.Answer) (*entity.Answer, error) {
	out, err := uc.answerRepo.Update(in)

	if err != nil {
		logger.Log.Errorf("AnswerApi.UpdateAnswer - r.Repo.UpdateAnswer: %v. Answer: %#+v", err, in)
	}

	return out, nil
}

func (uc *usecase) Delete(id int) error {
	return uc.answerRepo.Delete(id)
}
