package answer_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	crud_answers "gitlab.com/back1ng1/question-bot-api/app/usecase/crud_answer"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type restHandler struct {
	crud_answer_uc crud_answers.UseCase
}

func NewHandler(crud_answer_uc crud_answers.UseCase) RestHandler {
	return &restHandler{crud_answer_uc: crud_answer_uc}
}

func (h *restHandler) GetAnswer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("questionid"))
	if err != nil {
		return err
	}

	out, err := h.crud_answer_uc.Get(id)
	if err != nil {
		return err
	}

	return c.JSON(out)
}

func (h *restHandler) CreateAnswer(c *fiber.Ctx) error {
	answer := entity.Answer{}

	if err := c.BodyParser(&answer); err != nil {
		return err
	}

	out, err := h.crud_answer_uc.Create(answer)

	if err != nil {
		return err
	}

	return c.JSON(out)
}

func (h *restHandler) UpdateAnswer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return err
	}

	answer := entity.Answer{ID: int64(id)}
	if err := c.BodyParser(&answer); err != nil {
		return err
	}

	out, err := h.crud_answer_uc.Update(answer)
	if err != nil {
		return err
	}

	return c.JSON(out)
}

func (h *restHandler) DeleteAnswer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.JSON(ErrorResponse{Message: err.Error()})
	}

	return h.crud_answer_uc.Delete(id)
}
