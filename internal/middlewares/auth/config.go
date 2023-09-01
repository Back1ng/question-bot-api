package auth

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/app/repository"
)

type Config struct {
	Next func(*fiber.Ctx) bool

	// Filter defines a function to skip middleware.
	Filter func(c *fiber.Ctx) bool

	Repo repository.TokenRepository
}
