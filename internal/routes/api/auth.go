package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/pkg/tgauth"
)

type AuthApi struct {
	App  *fiber.App
	Repo database.AuthRepository
}

func (r *AuthApi) AuthRoutes() {
	r.App.Post("/api/auth/login", func(c *fiber.Ctx) error {
		var auth tgauth.Auth
		if err := c.BodyParser(&auth); err != nil {
			return err
		}

		// todo move generation to service
		token, err := r.Repo.GenerateToken(auth)

		if err != nil {
			return err
		}

		return c.JSON(token)
	})
}
