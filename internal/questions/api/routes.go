package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/questions/models"
)

func Routes(app *fiber.App) {
	app.Get("/api/questions", func(c *fiber.Ctx) error {
		questions := []models.Question{}

		database.GetConnection().Preload("Answers").Find(&questions)

		return c.JSON(questions)
	})

	app.Post("/api/questions", func(c *fiber.Ctx) error {
		// store question
		payload := models.Question{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		database.GetConnection().Create(&payload)

		return c.JSON(payload)
	})
}
