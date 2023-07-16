package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
)

func UserRoutes(app *fiber.App) {
	app.Get("/api/users", func(c *fiber.Ctx) error {
		users := []models.User{}

		database.Database.DB.Find(&users)

		return c.JSON(&users)
	})

	app.Post("/api/users/:id/preset", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))

		if err != nil {
			return err
		}

		user := models.User{}

		if err := c.BodyParser(&user); err != nil {
			return err
		}

		database.Database.DB.
			Preload("Preset").
			First(&user, models.User{Id: int64(id)}).
			UpdateColumn("preset_id", &user.Preset.ID)

		return c.JSON(&user)
	})
}
