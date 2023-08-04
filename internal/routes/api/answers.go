package api

import (
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type AnswerApi struct {
	App  *fiber.App
	Repo database.AnswerRepository
}

func (r *AnswerApi) AnswerRoutes() {
	r.App.Get("/api/answers/:questionid", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("questionid"))
		if err != nil {
			return err
		}

		answers, err := r.Repo.FindAnswersInQuestion(id)
		if err != nil {
			return err
		}

		if len(answers) == 0 {
			return c.JSON([]string{})
		}
		return c.JSON(answers)
	})
	// store new answer
	r.App.Post("/api/answer", func(c *fiber.Ctx) error {
		// store answer
		answer := entity.Answer{}

		if err := c.BodyParser(&answer); err != nil {
			return err
		}

		storedAnswer, err := r.Repo.StoreAnswer(answer)
		if err != nil {
			return err
		}

		return c.JSON(storedAnswer)
	})

	// update exists answer
	r.App.Put("/api/answer", func(c *fiber.Ctx) error {
		answer := entity.Answer{}
		if err := c.BodyParser(&answer); err != nil {
			return err
		}

		if err := r.Repo.UpdateAnswer(answer); err != nil {
			return err
		}

		return c.JSON(answer)
	})

	// delete exists answer by id
	r.App.Delete("/api/answer/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return err
		}

		err = r.Repo.DeleteAnswer(
			entity.Answer{ID: int64(id)},
		)
		if err != nil {
			return err
		}

		return c.JSON("success deleted")
	})
}
