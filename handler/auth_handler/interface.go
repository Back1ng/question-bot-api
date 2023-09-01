package auth_handler

import "github.com/gofiber/fiber/v2"

type AuthHandler interface {
	AuthLogin(c *fiber.Ctx) error
}
