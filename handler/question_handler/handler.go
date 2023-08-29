package question_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/app/usecase/crud_question"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type handler struct {
	crudQuestionUc crud_question.UseCase
}

func NewHandler(crudQuestionUc crud_question.UseCase, app *fiber.App) RestHandler {
	questionHandler := handler{crudQuestionUc: crudQuestionUc}

	app.Get("/api/questions/:presetid", questionHandler.GetByPreset)
	app.Post("/api/question", questionHandler.Create)
	app.Put("/api/question/:id", questionHandler.Update)
	app.Delete("/api/question/:id", questionHandler.Delete)

	return &questionHandler
}

func (h *handler) GetByPreset(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("presetid"))
	if err != nil {
		return err
	}

	questions, err := h.crudQuestionUc.GetByPreset(id)
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

	out, err := h.crudQuestionUc.Create(question)
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

	out, err := h.crudQuestionUc.Update(question)
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

	return h.crudQuestionUc.Delete(id)
}
