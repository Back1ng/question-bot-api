package answer_handler

import "github.com/gofiber/fiber/v2"

type RestHandler interface {
	GetAnswer(c *fiber.Ctx) error
	CreateAnswer(c *fiber.Ctx) error
	UpdateAnswer(c *fiber.Ctx) error
	DeleteAnswer(c *fiber.Ctx) error
}
