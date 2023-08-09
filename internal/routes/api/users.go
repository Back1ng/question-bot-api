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

func (r *UserApi) RegisterUserRoutes() {
	r.App.Get("/api/user/:chat_id", r.GetUser)
	r.App.Post("/api/user", r.StoreUser)
	r.App.Put("/api/user", r.UpdateUser)
}

func (r *UserApi) GetUser(c *fiber.Ctx) error {
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
}

func (r *UserApi) StoreUser(c *fiber.Ctx) error {
	var user entity.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	u, err := r.Repo.CreateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(u)
}

func (r *UserApi) UpdateUser(c *fiber.Ctx) error {
	user := entity.User{}
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	user, err := r.Repo.UpdateUser(user)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
