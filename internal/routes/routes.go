package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/routes/api"
	"gitlab.com/back1ng1/question-bot-api/internal/routes/web"
)

type SwaggerRoutes interface {
	RegisterSwaggerRoutes()
}

type QuestionRoutes interface {
	RegisterQuestionRoutes()
	GetQuestions(c *fiber.Ctx) error
	StoreQuestions(c *fiber.Ctx) error
	UpdateQuestion(c *fiber.Ctx) error
	DeleteQuestion(c *fiber.Ctx) error
}

type PresetRoutes interface {
	RegisterPresetRoutes()
	GetPresets(c *fiber.Ctx) error
	StorePreset(c *fiber.Ctx) error
	UpdatePreset(c *fiber.Ctx) error
	DeletePreset(c *fiber.Ctx) error
}

type AnswerRoutes interface {
	RegisterAnswerRoutes()
	GetAnswers(c *fiber.Ctx) error
	StoreAnswers(c *fiber.Ctx) error
	UpdateAnswer(c *fiber.Ctx) error
	DeleteAnswer(c *fiber.Ctx) error
}

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
	QuestionRoutes
	PresetRoutes
	AnswerRoutes
	UserRoutes
	AuthRoutes
	SwaggerRoutes
}

func RegisterRoutes(app *fiber.App, r database.Repositories) Routes {
	routes := Routes{
		QuestionRoutes: &api.QuestionApi{App: app, Repo: r.QuestionRepository},
		PresetRoutes:   &api.PresetApi{App: app, Repo: r.PresetRepository},
		UserRoutes:     &api.UserApi{App: app, Repo: r.UserRepository},
		AuthRoutes:     &api.AuthApi{App: app, Repo: r.AuthRepository},
		SwaggerRoutes:  &web.SwaggerApi{App: app},
	}

	routes.PresetRoutes.RegisterPresetRoutes()
	routes.QuestionRoutes.RegisterQuestionRoutes()
	routes.UserRoutes.RegisterUserRoutes()
	routes.AuthRoutes.RegisterAuthRoutes()
	routes.SwaggerRoutes.RegisterSwaggerRoutes()

	return routes
}
