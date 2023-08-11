package api

import (
	"strconv"

	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"

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
		logger.Log.Errorf("PresetApi.GetPresets - r.Repo.FindPresets: %v", err)
		return err
	}

	if len(presets) == 0 {
		logger.Log.Info("PresetApi.GetPresets - r.Repo.FindPresets: empty presets")
		return c.JSON([]string{})
	}

	return c.JSON(presets)
}

func (r *PresetApi) StorePreset(c *fiber.Ctx) error {
	preset := entity.Preset{}

	if err := c.BodyParser(&preset); err != nil {
		logger.Log.Errorf("PresetApi.StorePreset - c.BodyParser: %v", err)
		return err
	}

	p, err := r.Repo.StorePreset(preset)
	if err != nil {
		logger.Log.Errorf("PresetApi.StorePreset - r.Repo.StorePreset: %v", err)
		return err
	}

	return c.JSON(p)
}

func (r *PresetApi) UpdatePreset(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Errorf("PresetApi.UpdatePreset - strconv.Atoi: %v", err)
		return err
	}

	preset := entity.Preset{}
	if err := c.BodyParser(&preset); err != nil {
		logger.Log.Errorf("PresetApi.UpdatePreset - c.BodyParser: %v", err)
		return err
	}

	if err = r.Repo.UpdatePreset(id, preset); err != nil {
		logger.Log.Errorf("PresetApi.UpdatePreset - r.Repo.UpdatePreset: %v", err)
		return err
	}

	return c.JSON(preset)
}

func (r *PresetApi) DeletePreset(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Errorf("PresetApi.DeletePreset - strconv.Atoi: %v", err)
		return err
	}

	if err := r.Repo.DeletePreset(id); err != nil {
		logger.Log.Errorf("PresetApi.DeletePreset - r.Repo.DeletePreset: %v", err)
		return err
	}

	return c.JSON("success deleted")
}
