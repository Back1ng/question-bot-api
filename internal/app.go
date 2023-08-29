package app

import (
	"context"
	"fmt"
	"os"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	crud_answers "gitlab.com/back1ng1/question-bot-api/app/usecase/crud_answer"
	"gitlab.com/back1ng1/question-bot-api/app/usecase/crud_presets"
	"gitlab.com/back1ng1/question-bot-api/app/usecase/crud_question"
	"gitlab.com/back1ng1/question-bot-api/app/usecase/crud_user"
	"gitlab.com/back1ng1/question-bot-api/handler/answer_handler"
	"gitlab.com/back1ng1/question-bot-api/handler/preset_handler"
	"gitlab.com/back1ng1/question-bot-api/handler/question_handler"
	"gitlab.com/back1ng1/question-bot-api/handler/user_handler"
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/middlewares/auth"
	"gitlab.com/back1ng1/question-bot-api/internal/routes"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
	"gitlab.com/back1ng1/question-bot-api/repository/answer_repository_v1"
	"gitlab.com/back1ng1/question-bot-api/repository/preset_repository_v1"
	"gitlab.com/back1ng1/question-bot-api/repository/question_repository_v1"
	"gitlab.com/back1ng1/question-bot-api/repository/user_repository_v1"
)

func Run() {
	fmt.Println("Initializing configuration...")
	godotenv.Load(".env")

	fmt.Println("Initializing logging...")
	logger.New()

	fmt.Println("Initializing postgres connection...")
	sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		logger.Log.Errorf("app.Run - pgx.Connect: %v", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Println("Initializing repositories...")
	repo := database.GetRepositories(conn, sb)

	fmt.Println("Initializing api...")

	ignoreAuthPaths := []string{
		"/api/auth/login",
		"/swagger",
	}

	app := fiber.New()
	app.Use(cors.New())

	presets_repo := preset_repository_v1.NewRepository(conn, sb)
	crud_presets_uc := crud_presets.NewUseCase(presets_repo)
	preset_handler := preset_handler.NewHandler(crud_presets_uc)

	app.Get("/api/presets", preset_handler.GetAll)
	app.Post("/api/preset", preset_handler.Create)
	app.Put("/api/preset/:id", preset_handler.Update)
	app.Delete("/api/preset/:id", preset_handler.Delete)

	question_repo := question_repository_v1.NewRepository(conn, sb)
	crud_question_uc := crud_question.NewUseCase(question_repo)
	question_handler := question_handler.NewHandler(crud_question_uc)

	app.Get("/api/questions/:presetid", question_handler.GetByPreset)
	app.Post("/api/question", question_handler.Create)
	app.Put("/api/question/:id", question_handler.Update)
	app.Delete("/api/question/:id", question_handler.Delete)

	answer_repo := answer_repository_v1.NewRepository(conn, sb)
	crud_answer_uc := crud_answers.NewUseCase(answer_repo)
	answer_handler := answer_handler.NewHandler(crud_answer_uc)

	app.Get("/api/answers/:questionid", answer_handler.GetAnswer)
	app.Post("/api/answer", answer_handler.CreateAnswer)
	app.Put("/api/answer/:id", answer_handler.UpdateAnswer)
	app.Delete("/api/answer/:id", answer_handler.DeleteAnswer)

	user_repo := user_repository_v1.NewRepository(conn, sb)
	crud_user_uc := crud_user.NewUseCase(user_repo)
	user_handler := user_handler.NewHandler(crud_user_uc)

	app.Get("/api/user/interval/:id", user_handler.GetByInterval)
	app.Get("/api/user/:chat_id", user_handler.GetByChatId)
	app.Post("/api/user", user_handler.Create)
	app.Put("/api/user/:id", user_handler.Update)

	fmt.Println("Enable cors...")
	app.Use(auth.New(auth.Config{
		Repo: repo.AuthRepository,
		Filter: func(c *fiber.Ctx) bool {
			for _, v := range ignoreAuthPaths {
				if c.OriginalURL() == v {
					return true
				}

				if len(v) < len(c.OriginalURL()) && c.OriginalURL()[:len(v)] == v {
					return true
				}
			}

			return false
		},
	}))
	routes.RegisterRoutes(app, repo)

	app.Listen(":3000")
}
