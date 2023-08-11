package api

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
	"gitlab.com/back1ng1/question-bot-api/pkg/tgauth"
)

type AuthApi struct {
	App  *fiber.App
	Repo database.AuthRepository
}

func (r *AuthApi) RegisterAuthRoutes() {
	r.App.Post("/api/auth/login", r.AuthLogin)
}

func (r *AuthApi) AuthLogin(c *fiber.Ctx) error {
	var auth tgauth.Auth
	if err := c.BodyParser(&auth); err != nil {
		logger.Log.Errorf("AuthApi.AuthLogin - c.BodyParser(&auth): %v", err)
		return err
	}

	// todo move generation to service
	token, err := r.Repo.GenerateToken(auth)

	if err != nil {
		logger.Log.Errorf("AuthApi.AuthLogin - r.Repo.GenerateToken(auth): %v", err)
		return err
	}

	return c.JSON(token)
}
