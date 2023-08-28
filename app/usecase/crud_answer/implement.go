package crud_answers

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
)

type usecase struct {
	answer_repo repository.AnswerRepository
}

func NewUseCase(answer_repo repository.AnswerRepository) UseCase {
	return &usecase{answer_repo: answer_repo}
}

func (uc *usecase) Get(questionId int) ([]entity.Answer, error) {
	answers, err := uc.answer_repo.Get(questionId)

	if err != nil {
		logger.Log.Errorf("AnswerApi.GetAnswers - r.Repo.FindAnswersInQuestion: %v. QuestionId: %d", err, questionId)
	}

	return answers, nil
}

func (uc *usecase) Create(in entity.Answer) (entity.Answer, error) {
	out, err := uc.answer_repo.Create(in)

	if err != nil {
		logger.Log.Errorf("AnswerApi.StoreAnswers - r.Repo.StoreAnswer: %v. Answer: %#+v", err, in)
	}

	return out, nil
}

func (uc *usecase) Update(in entity.Answer) (entity.Answer, error) {
	out, err := uc.answer_repo.Update(in)

	if err != nil {
		logger.Log.Errorf("AnswerApi.UpdateAnswer - r.Repo.UpdateAnswer: %v. Answer: %#+v", err, in)
	}

	return out, nil
}

func (uc *usecase) Delete(id int) error {
	err := uc.answer_repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
