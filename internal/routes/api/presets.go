package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
)

func PresetRoutes(app *fiber.App) {
	app.Get("/api/presets", func(c *fiber.Ctx) error {
		preset := []models.Preset{}

		database.Database.DB.Find(&preset)

		return c.JSON(preset)
	})

	app.Get("/api/preset", func(c *fiber.Ctx) error {
		// send presets
		preset := models.Preset{}

		database.Database.DB.First(&preset)

		return c.JSON(preset)
	})

	app.Post("/api/preset", func(c *fiber.Ctx) error {
		// store preset
		preset := models.Preset{}

		if err := c.BodyParser(&preset); err != nil {
			return err
		}

		database.Database.DB.Create(&preset)

		return c.JSON(preset)
	})

	app.Put("/api/preset", func(c *fiber.Ctx) error {
		preset := models.Preset{}

		if err := c.BodyParser(&preset); err != nil {
			return err
		}

		database.Database.DB.
			First(&preset, models.Preset{Id: preset.Id}).
			Updates(&preset)

		return c.JSON(preset)
	})

	app.Delete("/api/preset/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return err
		}

		preset := models.Preset{Id: int64(id)}

		database.Database.DB.Delete(&preset, models.Preset{Id: preset.Id})

		return c.JSON(preset)
	})
}
