package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/routes/api"
)

type AuthRoutes interface {
	RegisterAuthRoutes()
	AuthLogin(c *fiber.Ctx) error
}

type Routes struct {
	AuthRoutes
}

func RegisterRoutes(app *fiber.App, r database.Repositories) Routes {
	routes := Routes{
		AuthRoutes: &api.AuthApi{App: app, Repo: r.AuthRepository},
	}

	routes.AuthRoutes.RegisterAuthRoutes()

	return routes
}
