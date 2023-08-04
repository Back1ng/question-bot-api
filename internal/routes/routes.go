package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/routes/api"
)

type QuestionRoutes interface {
	QuestionRoutes()
}

type PresetRoutes interface {
	PresetRoutes()
}

type AnswerRoutes interface {
	AnswerRoutes()
}

type UserRoutes interface {
	UserRoutes()
}

type AuthRoutes interface {
	AuthRoutes()
}

type Routes struct {
	QuestionRoutes
	PresetRoutes
	AnswerRoutes
	UserRoutes
	AuthRoutes
}

func RegisterRoutes(app *fiber.App, r database.Repositories) Routes {
	routes := Routes{
		QuestionRoutes: &api.QuestionApi{App: app, Repo: r.QuestionRepository},
		PresetRoutes:   &api.PresetApi{App: app, Repo: r.PresetRepository},
		AnswerRoutes:   &api.AnswerApi{App: app, Repo: r.AnswerRepository},
		UserRoutes:     &api.UserApi{App: app, Repo: r.UserRepository},
		AuthRoutes:     &api.AuthApi{App: app, Repo: r.AuthRepository},
	}

	routes.PresetRoutes.PresetRoutes()
	routes.AnswerRoutes.AnswerRoutes()
	routes.QuestionRoutes.QuestionRoutes()
	routes.UserRoutes.UserRoutes()
	routes.AuthRoutes.AuthRoutes()

	return routes
}
