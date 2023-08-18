package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "gitlab.com/back1ng1/question-bot-api/cmd/app/docs"
)

type SwaggerApi struct {
	App *fiber.App
}

func (r *SwaggerApi) RegisterSwaggerRoutes() {
	r.App.Get("/swagger/*", swagger.HandlerDefault)
}
