package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/routes/api"
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

type Routes struct {
	QuestionRoutes
	PresetRoutes
	AnswerRoutes
	UserRoutes
}

func RegisterRoutes(app *fiber.App, r database.Repositories) Routes {
	routes := Routes{
		QuestionRoutes: &api.QuestionApi{App: app, Repo: r.QuestionRepository},
		PresetRoutes:   &api.PresetApi{App: app, Repo: r.PresetRepository},
		AnswerRoutes:   &api.AnswerApi{App: app, Repo: r.AnswerRepository},
		UserRoutes:     &api.UserApi{App: app, Repo: r.UserRepository},
	}

	routes.PresetRoutes.PresetRoutes()
	routes.AnswerRoutes.AnswerRoutes()
	routes.QuestionRoutes.QuestionRoutes()
	routes.UserRoutes.UserRoutes()

	return routes
}
