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

func (r *PresetApi) RegisterPresetRoutes() {
	r.App.Get("/api/presets", r.GetPresets)
	r.App.Post("/api/preset", r.StorePreset)
	r.App.Put("/api/preset/:id", r.UpdatePreset)
	r.App.Delete("/api/preset/:id", r.DeletePreset)
}

func (r *PresetApi) GetPresets(c *fiber.Ctx) error {
	presets, err := r.Repo.FindPresets()

	if err != nil {
		return err
	}

	if len(presets) == 0 {
		return c.JSON([]string{})
	}

	return c.JSON(presets)
}

func (r *PresetApi) StorePreset(c *fiber.Ctx) error {
	preset := entity.Preset{}

	if err := c.BodyParser(&preset); err != nil {
		return err
	}

	p, err := r.Repo.StorePreset(preset)
	if err != nil {
		return err
	}

	return c.JSON(p)
}

func (r *PresetApi) UpdatePreset(c *fiber.Ctx) error {
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
}

func (r *PresetApi) DeletePreset(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return err
	}

	if err := r.Repo.DeletePreset(id); err != nil {
		return err
	}

	return c.JSON("success deleted")
}
