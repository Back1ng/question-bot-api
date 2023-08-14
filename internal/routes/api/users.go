package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
)

type UserApi struct {
	App  *fiber.App
	Repo database.UserRepository
}

func (r *UserApi) RegisterUserRoutes() {
	r.App.Get("/api/user/interval/:id", r.GetUsersByInterval)
	r.App.Get("/api/user/:chat_id", r.GetUser)
	r.App.Post("/api/user", r.StoreUser)
	r.App.Put("/api/user/:id", r.UpdateUser)
}

func (r *UserApi) GetUsersByInterval(c *fiber.Ctx) error {
	intervalId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logger.Log.Errorf("UserApi.GetUsersByInterval - strconv.Atoi: %v", err)
		return err
	}

	users, err := r.Repo.FindUsersByInterval(intervalId)
	if err != nil {
		logger.Log.Errorf("UserApi.GetUsersByInterval - r.Repo.FindUsersByInterval: %v", err)
		return err
	}

	if len(users) == 0 {
		return c.JSON([]string{})
	}

	return c.JSON(users)
}

func (r *UserApi) GetUser(c *fiber.Ctx) error {
	chatId, err := strconv.Atoi(c.Params("chat_id"))
	if err != nil {
		logger.Log.Errorf("UserApi.GetUser - strconv.Atoi: %v", err)
		return err
	}

	user, err := r.Repo.FindUserByChatId(chatId)
	if err != nil {
		logger.Log.Errorf("UserApi.GetUser - r.Repo.FindUserByChatId: %v", err)
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
		logger.Log.Errorf("UserApi.StoreUser - c.BodyParser: %v", err)
		return err
	}

	u, err := r.Repo.CreateUser(user)
	if err != nil {
		logger.Log.Errorf("UserApi.StoreUser - r.Repo.CreateUser: %v", err)
		return err
	}

	return c.JSON(u)
}

func (r *UserApi) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		logger.Log.Errorf("UserApi.UpdateUser - strconv.Atoi: %v", err)
		return err
	}

	user := entity.User{ID: int64(id)}
	if err := c.BodyParser(&user); err != nil {
		logger.Log.Errorf("UserApi.UpdateUser - c.BodyParser: %v", err)
		return err
	}

	user, err = r.Repo.UpdateUser(user)
	if err != nil {
		logger.Log.Errorf("UserApi.UpdateUser - r.Repo.UpdateUser: %v", err)
		return err
	}

	return c.JSON(user)
}
