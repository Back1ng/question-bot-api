package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
)

func AnswerRoutes(app *fiber.App) {
	app.Get("/api/answers/:questionid", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("questionid"))

		if err != nil {
			return err
		}

		answers := []models.Answer{}

		database.Database.DB.
			Preload("Question").
			Find(&answers, &models.Answer{QuestionId: int64(id)})

		return c.JSON(answers)
	})

	// send info about answer
	app.Get("/api/answer", func(c *fiber.Ctx) error {
		answer := models.Answer{}

		database.Database.DB.First(&answer)

		return c.JSON(answer)
	})

	// store new answer
	app.Post("/api/answer", func(c *fiber.Ctx) error {
		// store answer
		answer := models.Answer{}

		if err := c.BodyParser(&answer); err != nil {
			return err
		}

		database.Database.DB.Create(&answer)

		return c.JSON(answer)
	})

	// update exists answer
	app.Put("/api/answer", func(c *fiber.Ctx) error {
		answer := models.Answer{}

		if err := c.BodyParser(&answer); err != nil {
			return err
		}

		dbAnswer := models.Answer{}
		database.Database.DB.
			First(&dbAnswer, models.Answer{ID: answer.ID}).
			Updates(&answer)

		return c.JSON(answer)
	})

	// delete exists answer by id
	app.Delete("/api/answer/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return err
		}

		answer := models.Answer{}

		database.Database.DB.Delete(&answer, &models.Answer{ID: int64(id)})

		return c.JSON(answer)
	})
}
