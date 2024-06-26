package question_handler

import "github.com/gofiber/fiber/v2"

type RestHandler interface {
	GetByPreset(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}
