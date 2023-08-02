package api

import (
	"gitlab.com/back1ng1/question-bot/internal/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database/entity"
)

type QuestionApi struct {
	App *fiber.App
	database.QuestionRepository
}

func (r *QuestionApi) QuestionRoutes() {
	r.App.Get("/api/questions/:presetid", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("presetid"))
		if err != nil {
			return err
		}

		questions, err := r.FindQuestionsInPreset(id)
		if err != nil {
			return err
		}

		if len(questions) == 0 {
			return c.JSON([]string{})
		}
		return c.JSON(questions)
	})

	// store new question with all needed data
	r.App.Post("/api/question", func(c *fiber.Ctx) error {
		var question entity.Question
		if err := c.BodyParser(&question); err != nil {
			return err
		}

		if err := r.StoreQuestion(question); err != nil {
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

		if err = r.UpdateQuestionTitle(id, question); err != nil {
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

		if err = r.DeleteQuestion(id); err != nil {
			return err
		}

		return c.JSON("success deleted")
	})
}
