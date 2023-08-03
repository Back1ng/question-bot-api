package api

import "C"
import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"strconv"
)

type UserApi struct {
	App  *fiber.App
	Repo database.UserRepository
}

func (r *UserApi) UserRoutes() {
	r.App.Get("/api/user/:chat_id", func(c *fiber.Ctx) error {
		chatId, err := strconv.Atoi(c.Params("chat_id"))
		if err != nil {
			return err
		}

		user, err := r.Repo.FindUserByChatId(chatId)
		if err != nil {
			return err
		}

		if user.ID == 0 {
			return c.JSON([]string{})
		}
		return c.JSON(user)
	})
	r.App.Post("/api/user", func(c *fiber.Ctx) error {
		var user entity.User
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		u, err := r.Repo.CreateUser(user)
		if err != nil {
			return err
		}

		return c.JSON(u)
	})
	r.App.Put("/api/user", func(c *fiber.Ctx) error {
		user := entity.User{}
		if err := c.BodyParser(&user); err != nil {
			return err
		}

		user, err := r.Repo.UpdateUser(user)
		if err != nil {
			return err
		}

		return c.JSON(user)
	})
	/*
		app.Get("/api/users", func(c *fiber.Ctx) error {
			users := []entity.User{}

			database.Database.DB.Find(&users)

			return c.JSON(&users)
		})

		app.Post("/api/users/:id/preset", func(c *fiber.Ctx) error {
			id, err := strconv.Atoi(c.Params("id"))

			if err != nil {
				return err
			}

			user := entity.User{}

			if err := c.BodyParser(&user); err != nil {
				return err
			}

			preset := user.PresetId

			user.ID = int64(id)
			first_user := database.Database.DB.First(&user)

			user.PresetId = preset
			first_user.Updates(&user)

			return c.JSON(&user)
		})


	*/
}
