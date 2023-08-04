package app

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"gitlab.com/back1ng1/question-bot-api/internal/database"
	"gitlab.com/back1ng1/question-bot-api/internal/middlewares/auth"
	"gitlab.com/back1ng1/question-bot-api/internal/routes"
	"os"
)

func Run() {
	fmt.Println("Initializing configuration...")
	godotenv.Load(".env")

	fmt.Println("Initializing postgres connection...")
	sb := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Println("Initializing repositories...")
	repo := database.GetRepositories(conn, sb)

	fmt.Println("Initializing api...")

	ignoreAuthPaths := []string{
		"/api/auth/login",
	}

	app := fiber.New()
	app.Use(cors.New())
	app.Use(auth.New(auth.Config{
		Repo: repo.AuthRepository,
		Filter: func(c *fiber.Ctx) bool {
			for _, v := range ignoreAuthPaths {
				if c.OriginalURL() == v {
					return true
				}
			}

			return false
		},
	}))
	routes.RegisterRoutes(app, repo)

	app.Listen(":3000")
}
