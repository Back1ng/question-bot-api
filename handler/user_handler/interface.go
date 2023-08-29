package user_handler

import "github.com/gofiber/fiber/v2"

type RestHandler interface {
	GetByInterval(c *fiber.Ctx) error
	GetByChatId(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
}
