package api

import (
	"gitlab.com/back1ng1/question-bot/internal/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database/entity"
)

type QuestionApi struct {
	App  *fiber.App
	Repo database.QuestionRepository
}

func (r *QuestionApi) QuestionRoutes() {
	r.App.Get("/api/questions/:presetid", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("presetid"))
		if err != nil {
			return err
		}

		questions, err := r.Repo.FindQuestionsInPreset(id)
		if err != nil {
			return err
		}

		return c.JSON(questions)
	})

	// store new question with all needed data
	r.App.Post("/api/question", func(c *fiber.Ctx) error {
		var question entity.Question
		if err := c.BodyParser(&question); err != nil {
			return err
		}

		if err := r.Repo.StoreQuestion(question); err != nil {
			return err
		}

		return c.JSON(question)
	})

	// update title in existence question by id
	r.App.Put("/api/question/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		var question entity.Question
		if err := c.BodyParser(&question); err != nil {
			return err
		}

		if err = r.Repo.UpdateQuestionTitle(id, question); err != nil {
			return err
		}

		return c.JSON(question)
	})

	// delete question by him Id
	r.App.Delete("/api/question/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		if err = r.Repo.DeleteQuestion(id); err != nil {
			return err
		}

		return c.JSON("success deleted")
	})
}
