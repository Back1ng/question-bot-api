package api

import (
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PresetApi struct {
	App  *fiber.App
	Repo database.PresetRepository
}

func (r *PresetApi) PresetRoutes() {
	r.App.Get("/api/presets", func(c *fiber.Ctx) error {
		presets, err := r.Repo.FindPresets()

		if err != nil {
			return err
		}

		if len(presets) == 0 {
			return c.JSON([]string{})
		}
		return c.JSON(presets)
	})
	r.App.Post("/api/preset", func(c *fiber.Ctx) error {
		preset := entity.Preset{}

		if err := c.BodyParser(&preset); err != nil {
			return err
		}

		if err := r.Repo.StorePreset(preset); err != nil {
			return err
		}

		return c.JSON(preset)
	})
	r.App.Put("/api/preset/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		preset := entity.Preset{}
		if err := c.BodyParser(&preset); err != nil {
			return err
		}

		if err = r.Repo.UpdatePreset(id, preset); err != nil {
			return err
		}

		return c.JSON(preset)
	})
	r.App.Delete("/api/preset/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return err
		}

		if err := r.Repo.DeletePreset(id); err != nil {
			return err
		}

		return c.JSON("success deleted")
	})
}
