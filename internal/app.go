package app

import (
	"context"
	"fmt"
	"os"

	"gitlab.com/back1ng1/question-bot-api/app/usecase/auth_usecase"
	"gitlab.com/back1ng1/question-bot-api/handler/auth_handler"
	"gitlab.com/back1ng1/question-bot-api/repository/tokens_repository_v1"

	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"
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
	"gitlab.com/back1ng1/question-bot-api/internal/middlewares/auth"
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
	conn, err := pgx.Connect(context.Background(), os.Getenv("PGBOUNCER_URL"))
	if err != nil {
		logger.Log.Errorf("app.Run - pgx.Connect: %v", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Println("Initializing api...")

	ignoreAuthPaths := []string{
		"/api/auth/login",
	}

	app := fiber.New()

	fmt.Println("Initializing fiber logging...")
	app.Use(fiber_logger.New())

	fmt.Println("Initializing cors...")
	app.Use(cors.New())

	fmt.Println("Initializing presets api...")
	presetsRepo := preset_repository_v1.NewRepository(conn, sb)
	crudPresetsUc := crud_presets.NewUseCase(presetsRepo)
	preset_handler.NewHandler(crudPresetsUc, app)

	fmt.Println("Initializing questions api...")
	questionRepo := question_repository_v1.NewRepository(conn, sb)
	crudQuestionUc := crud_question.NewUseCase(questionRepo)
	question_handler.NewHandler(crudQuestionUc, app)

	fmt.Println("Initializing answers api...")
	answerRepo := answer_repository_v1.NewRepository(conn, sb)
	crudAnswerUc := crud_answers.NewUseCase(answerRepo)
	answer_handler.NewHandler(crudAnswerUc, app)

	fmt.Println("Initializing users api...")
	userRepo := user_repository_v1.NewRepository(conn, sb)
	crudUserUc := crud_user.NewUseCase(userRepo)
	user_handler.NewHandler(crudUserUc, app)

	fmt.Println("Initializing auth api...")
	tokensRepo := tokens_repository_v1.NewRepository(conn, sb)
	authUseCase := auth_usecase.NewUsecase(tokensRepo)
	auth_handler.NewHandler(authUseCase, app)

	app.Use(auth.New(auth.Config{
		Repo: tokensRepo,
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

	app.Listen(":3000")
}
