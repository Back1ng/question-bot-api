package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database/entity"
	"gitlab.com/back1ng1/question-bot/internal/database/repository"
)

func PresetRoutes(app *fiber.App) {
	app.Get("/api/presets", func(c *fiber.Ctx) error {
		presets, err := repository.FindPresets()

		if err != nil {
			return err
		}

		return c.JSON(presets)
	})
	app.Post("/api/preset", func(c *fiber.Ctx) error {
		preset := entity.Preset{}

		if err := c.BodyParser(&preset); err != nil {
			return err
		}

		if err := repository.StorePreset(preset); err != nil {
			return err
		}

		return c.JSON(preset)
	})
	app.Put("/api/preset/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		preset := entity.Preset{}
		if err := c.BodyParser(&preset); err != nil {
			return err
		}

		if err = repository.UpdatePreset(id, preset); err != nil {
			return err
		}

		return c.JSON(preset)
	})
	app.Delete("/api/preset/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return err
		}

		if err := repository.DeletePreset(id); err != nil {
			return err
		}

		return c.JSON("success deleted")
	})
}
