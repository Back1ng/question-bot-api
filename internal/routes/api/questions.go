package api

import (
	"strconv"

	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type QuestionApi struct {
	App  *fiber.App
	Repo database.QuestionRepository
}

func (r *QuestionApi) RegisterQuestionRoutes() {
	r.App.Get("/api/questions/:presetid", r.GetQuestions)
	r.App.Post("/api/question", r.StoreQuestions)
	r.App.Put("/api/question/:id", r.UpdateQuestion)
	r.App.Delete("/api/question/:id", r.DeleteQuestion)
}

func (r *QuestionApi) GetQuestions(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("presetid"))
	if err != nil {
		logger.Log.Errorf("QuestionApi.GetQuestions - strconv.Atoi: %v", err)
		return err
	}

	questions, err := r.Repo.FindQuestionsInPreset(id)
	if err != nil {
		logger.Log.Errorf("QuestionApi.GetQuestions - r.Repo.FindQuestionsInPreset: %v", err)
		return err
	}

	if len(questions) == 0 {
		logger.Log.Info("QuestionApi.GetQuestions - r.Repo.FindQuestionsInPreset: empty questions")
		return c.JSON([]string{})
	}
	return c.JSON(questions)
}

func (r *QuestionApi) StoreQuestions(c *fiber.Ctx) error {
	var question entity.Question
	if err := c.BodyParser(&question); err != nil {
		logger.Log.Errorf("QuestionApi.StoreQuestions - c.BodyParser: %v", err)
		return err
	}

	if err := r.Repo.StoreQuestion(question); err != nil {
		logger.Log.Errorf("QuestionApi.StoreQuestions - r.Repo.StoreQuestion: %v", err)
		return err
	}

	return c.JSON(question)
}

func (r *QuestionApi) UpdateQuestion(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Errorf("QuestionApi.UpdateQuestion - strconv.Atoi: %v", err)
		return err
	}

	var question entity.Question
	if err := c.BodyParser(&question); err != nil {
		logger.Log.Errorf("QuestionApi.UpdateQuestion - c.BodyParser: %v", err)
		return err
	}

	if err = r.Repo.UpdateQuestionTitle(id, question); err != nil {
		logger.Log.Errorf("QuestionApi.UpdateQuestion - r.Repo.UpdateQuestionTitle: %v", err)
		return err
	}

	return c.JSON(question)
}

func (r *QuestionApi) DeleteQuestion(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Errorf("QuestionApi.DeleteQuestion - strconv.Atoi: %v", err)
		return err
	}

	if err = r.Repo.DeleteQuestion(id); err != nil {
		logger.Log.Errorf("QuestionApi.DeleteQuestion - r.Repo.DeleteQuestion: %v", err)
		return err
	}

	return c.JSON("success deleted")
}
