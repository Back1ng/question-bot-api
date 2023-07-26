package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
	"gitlab.com/back1ng1/question-bot/internal/database/repository"
)

func QuestionRoutes(app *fiber.App) {
	app.Get("/api/questions/:presetid", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("presetid"))
		if err != nil {
			return err
		}

		questions, err := repository.FindQuestionsInPreset(id)
		if err != nil {
			return err
		}

		return c.JSON(questions)
	})

	// store new question with all needed data
	app.Post("/api/question", func(c *fiber.Ctx) error {
		question := models.Question{}

		if err := c.BodyParser(&question); err != nil {
			return err
		}

		question, err := repository.StoreQuestion(question)

		if err != nil {
			return err
		}

		return c.JSON(question)
	})

	// update title in existence question by id
	app.Put("/api/question/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		question := models.Question{}
		if err := c.BodyParser(&question); err != nil {
			return err
		}

		_, err = repository.UpdateQuestionTitle(id, question)
		if err != nil {
			return err
		}

		return c.JSON(question)
	})

	// delete question by him Id
	app.Delete("/api/question/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		err = repository.DeleteQuestion(id)
		if err != nil {
			return err
		}

		return c.JSON("success deleted")
	})
}
