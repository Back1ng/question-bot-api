package app

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/routes"
	"os"
)

func Run() {
	fmt.Println("Initializing configuration...")
	godotenv.Load(".env")

	fmt.Println("Initializing postgres connection...")
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRESQL_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	fmt.Println("Initializing repositories...")
	repo := database.GetRepositories(conn)

	fmt.Println("Initializing api...")
	app := fiber.New()
	app.Use(cors.New())
	routes.RegisterRoutes(app, repo)

	app.Listen(":3000")
}
