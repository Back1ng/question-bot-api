package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
)

func QuestionRoutes(app *fiber.App) {
	// get question by id
	app.Get("/api/question", func(c *fiber.Ctx) error {
		question := models.Question{}

		database.Database.DB.Preload("Answers").Find(&question)

		return c.JSON(question)
	})

	// store new question with all needed data
	app.Post("/api/question", func(c *fiber.Ctx) error {
		payload := models.Question{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		database.Database.DB.Create(&payload)

		return c.JSON(payload)
	})

	// update exists question by him ID with needed data
	app.Put("/api/question", func(c *fiber.Ctx) error {
		payload := models.Question{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		if payload.ID == 0 {
			return errors.New("ID not represented")
		}

		database.Database.DB.
			First(&payload, models.Question{ID: payload.ID}).
			Updates(&payload)

		return c.JSON(payload)
	})

	// delete question by him Id
	app.Delete("/api/question", func(c *fiber.Ctx) error {
		payload := models.Question{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		database.Database.DB.Delete(payload, models.Question{ID: payload.ID})

		return c.JSON(payload)
	})
}
