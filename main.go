package main

import (
	"context"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/routes/api"
)

func main() {
	godotenv.Load(".env")

	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	database.Database = database.DbInstance{DB: conn}

	app := fiber.New()

	app.Use(cors.New())

	// register routes
	api.QuestionRoutes(app)
	api.PresetRoutes(app)
	api.AnswerRoutes(app)
	// api.UserRoutes(app)

	app.Listen(":3000")
}
