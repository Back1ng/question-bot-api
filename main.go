package main

import (
	"context"
	"fmt"
	"gitlab.com/back1ng1/question-bot/internal/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"gitlab.com/back1ng1/question-bot/internal/database"
)

func main() {
	godotenv.Load(".env")

	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	repo := database.GetRepositories(conn)

	app := fiber.New()
	app.Use(cors.New())

	routes := routes.RegisterRoutes(app, repo)
	routes.PresetRoutes.PresetRoutes()
	routes.AnswerRoutes.AnswerRoutes()
	routes.QuestionRoutes.QuestionRoutes()
	routes.UserRoutes.UserRoutes()

	app.Listen(":3000")
}
