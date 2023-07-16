package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
)

func AnswerRoutes(app *fiber.App) {
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

		database.Database.DB.
			First(&answer, models.Answer{ID: answer.ID}).
			Updates(&answer)

		return c.JSON(answer)
	})

	// delete exists answer by id
	app.Delete("/api/answer", func(c *fiber.Ctx) error {
		answer := models.Answer{}

		if err := c.BodyParser(&answer); err != nil {
			return err
		}

		database.Database.DB.Delete(&answer, models.Answer{ID: answer.ID})

		return c.JSON(answer)
	})
}
