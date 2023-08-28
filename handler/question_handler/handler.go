package question_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/app/usecase/crud_question"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type handler struct {
	crud_question_uc crud_question.UseCase
}

func NewHandler(crud_question_uc crud_question.UseCase) RestHandler {
	return &handler{crud_question_uc: crud_question_uc}
}

func (h *handler) GetByPreset(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("presetid"))
	if err != nil {
		return err
	}

	questions, err := h.crud_question_uc.GetByPreset(id)
	if err != nil {
		return err
	}

	return c.JSON(questions)
}

func (h *handler) Create(c *fiber.Ctx) error {
	var question entity.Question
	if err := c.BodyParser(&question); err != nil {
		return err
	}

	out, err := h.crud_question_uc.Create(question)
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

	question := entity.Question{ID: int64(id)}
	if err := c.BodyParser(&question); err != nil {
		return err
	}

	out, err := h.crud_question_uc.Update(question)
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

	return h.crud_question_uc.Delete(id)
}
