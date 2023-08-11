package api

import (
	"strconv"

	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type AnswerApi struct {
	App  *fiber.App
	Repo database.AnswerRepository
}

func (r *AnswerApi) RegisterAnswerRoutes() {
	r.App.Get("/api/answers/:questionid", r.GetAnswers)
	r.App.Post("/api/answer", r.StoreAnswers)
	r.App.Put("/api/answer/:id", r.UpdateAnswer)
	r.App.Delete("/api/answer/:id", r.DeleteAnswer)
}

func (r *AnswerApi) GetAnswers(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("questionid"))
	if err != nil {
		logger.Log.Errorf("AnswerApi.GetAnswers - strconv.Atoi: %v", err)
		return err
	}

	answers, err := r.Repo.FindAnswersInQuestion(id)
	if err != nil {
		logger.Log.Errorf("AnswerApi.GetAnswers - r.Repo.FindAnswersInQuestion: %v", err)
		return err
	}

	if len(answers) == 0 {
		logger.Log.Info("AnswerApi.GetAnswers - r.Repo.FindAnswersInQuestion: empty answers")
		return c.JSON([]string{})
	}
	return c.JSON(answers)
}

func (r *AnswerApi) StoreAnswers(c *fiber.Ctx) error {
	answer := entity.Answer{}

	if err := c.BodyParser(&answer); err != nil {
		logger.Log.Errorf("AnswerApi.StoreAnswers - c.BodyParser: %v", err)
		return err
	}

	storedAnswer, err := r.Repo.StoreAnswer(answer)
	if err != nil {
		logger.Log.Errorf("AnswerApi.StoreAnswers - r.Repo.StoreAnswer: %v", err)
		return err
	}

	return c.JSON(storedAnswer)
}

func (r *AnswerApi) UpdateAnswer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Errorf("AnswerApi.UpdateAnswer - strconv.Atoi: %v", err)
		return err
	}

	answer := entity.Answer{ID: int64(id)}
	if err := c.BodyParser(&answer); err != nil {
		logger.Log.Errorf("AnswerApi.UpdateAnswer - c.BodyParser: %v", err)
		return err
	}

	if err := r.Repo.UpdateAnswer(answer); err != nil {
		logger.Log.Errorf("AnswerApi.UpdateAnswer - r.Repo.UpdateAnswer: %v", err)
		return err
	}

	return c.JSON(answer)
}

func (r *AnswerApi) DeleteAnswer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Errorf("AnswerApi.DeleteAnswer - strconv.Atoi: %v", err)
		return err
	}

	err = r.Repo.DeleteAnswer(
		entity.Answer{ID: int64(id)},
	)
	if err != nil {
		logger.Log.Errorf("AnswerApi.DeleteAnswer - r.Repo.DeleteAnswer: %v", err)
		return err
	}

	return c.JSON("success deleted")
}
