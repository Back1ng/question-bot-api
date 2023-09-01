package auth_handler

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/app/usecase/auth_usecase"
	"gitlab.com/back1ng1/question-bot-api/pkg/tgauth"
)

type handler struct {
	usecase auth_usecase.UseCase
	app     *fiber.App
}

func NewHandler(usecase auth_usecase.UseCase, app *fiber.App) AuthHandler {
	h := handler{
		usecase: usecase,
		app:     app,
	}

	app.Post("/api/auth/login", h.AuthLogin)

	return &h
}

func (h *handler) AuthLogin(c *fiber.Ctx) error {
	var auth tgauth.Auth
	if err := c.BodyParser(&auth); err != nil {
		return err
	}

	hash, err := h.usecase.Authenticate(auth)
	if err != nil {
		return err
	}

	return c.JSON(hash)
}
