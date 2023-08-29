package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/routes/api"
)

type UserRoutes interface {
	RegisterUserRoutes()
	GetUsersByInterval(c *fiber.Ctx) error
	GetUser(c *fiber.Ctx) error
	StoreUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
}

type AuthRoutes interface {
	RegisterAuthRoutes()
	AuthLogin(c *fiber.Ctx) error
}

type Routes struct {
	UserRoutes
	AuthRoutes
}

func RegisterRoutes(app *fiber.App, r database.Repositories) Routes {
	routes := Routes{
		UserRoutes: &api.UserApi{App: app, Repo: r.UserRepository},
		AuthRoutes: &api.AuthApi{App: app, Repo: r.AuthRepository},
	}

	routes.UserRoutes.RegisterUserRoutes()
	routes.AuthRoutes.RegisterAuthRoutes()

	return routes
}
