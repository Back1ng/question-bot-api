package preset_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/app/usecase/crud_presets"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type handler struct {
	crud_presets_uc crud_presets.UseCase
}

func NewHandler(crud_presets_uc crud_presets.UseCase) RestHandler {
	return &handler{crud_presets_uc: crud_presets_uc}
}

func (h *handler) GetAll(c *fiber.Ctx) error {
	out, err := h.crud_presets_uc.GetAll()
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

	out, err := h.crud_presets_uc.Create(preset)
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

	out, err := h.crud_presets_uc.Update(preset)
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

	err = h.crud_presets_uc.Delete(int64(id))
	if err != nil {
		return err
	}

	return c.JSON(err)
}
