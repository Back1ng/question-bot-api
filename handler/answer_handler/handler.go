package answer_handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	crud_answers "gitlab.com/back1ng1/question-bot-api/app/usecase/crud_answer"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type restHandler struct {
	crudAnswerUc crud_answers.UseCase
}

func NewHandler(crudAnswerUc crud_answers.UseCase, app *fiber.App) RestHandler {
	answerHandler := restHandler{crudAnswerUc: crudAnswerUc}

	app.Get("/api/answers/:questionid", answerHandler.GetAnswer)
	app.Post("/api/answer", answerHandler.CreateAnswer)
	app.Put("/api/answer/:id", answerHandler.UpdateAnswer)
	app.Delete("/api/answer/:id", answerHandler.DeleteAnswer)

	return &answerHandler
}

func (h *restHandler) GetAnswer(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("questionid"))
	if err != nil {
		return err
	}

	out, err := h.crudAnswerUc.Get(id)
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

	out, err := h.crudAnswerUc.Create(answer)
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

	out, err := h.crudAnswerUc.Update(answer)
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

	return h.crudAnswerUc.Delete(id)
}
