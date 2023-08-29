package preset_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/app/usecase/crud_presets"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type handler struct {
	crudPresetsUc crud_presets.UseCase
}

func NewHandler(crudPresetsUc crud_presets.UseCase) RestHandler {
	return &handler{crudPresetsUc: crudPresetsUc}
}

func (h *handler) GetAll(c *fiber.Ctx) error {
	out, err := h.crudPresetsUc.GetAll()
	if err != nil {
		return err
	}

	if len(out) == 0 {
		return c.JSON([]string{})
	}

	return c.JSON(out)
}

func (h *handler) Create(c *fiber.Ctx) error {
	preset := entity.Preset{}

	if err := c.BodyParser(&preset); err != nil {
		return err
	}

	out, err := h.crudPresetsUc.Create(preset)
	if err != nil {
		return err
	}

	return c.JSON(out)
}

func (h *handler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	preset := entity.Preset{ID: int64(id)}
	if err := c.BodyParser(&preset); err != nil {
		return err
	}

	out, err := h.crudPresetsUc.Update(preset)
	if err != nil {
		return err
	}

	return c.JSON(out)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	err = h.crudPresetsUc.Delete(int64(id))
	if err != nil {
		return err
	}

	return c.JSON(err)
}
