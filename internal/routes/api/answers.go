package api

import (
	"gitlab.com/back1ng1/question-bot/internal/database"
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

		return c.JSON(answers)
	})
	/*

		// send info about answer
		app.Get("/api/answer", func(c *fiber.Ctx) error {
			answer := entity.Answer{}

			database.Database.DB.First(&answer)

			return c.JSON(answer)
		})

		// store new answer
		app.Post("/api/answer", func(c *fiber.Ctx) error {
			// store answer
			answer := entity.Answer{}

			if err := c.BodyParser(&answer); err != nil {
				return err
			}

			database.Database.DB.Create(&answer)

			return c.JSON(answer)
		})

		// update exists answer
		app.Put("/api/answer", func(c *fiber.Ctx) error {
			answer := entity.Answer{}

			if err := c.BodyParser(&answer); err != nil {
				return err
			}

			dbAnswer := entity.Answer{}
			database.Database.DB.
				First(&dbAnswer, entity.Answer{ID: answer.ID}).
				Updates(&answer)

			return c.JSON(answer)
		})

		// delete exists answer by id
		app.Delete("/api/answer/:id", func(c *fiber.Ctx) error {
			id, err := strconv.Atoi(c.Params("id"))

			if err != nil {
				return err
			}

			answer := entity.Answer{}

			database.Database.DB.Delete(&answer, &entity.Answer{ID: int64(id)})

			return c.JSON(answer)
		})
	*/
}
