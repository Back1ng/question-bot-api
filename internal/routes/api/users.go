package api

import (
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
	/*
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

			preset := user.PresetId

			user.ID = int64(id)
			first_user := database.Database.DB.First(&user)

			user.PresetId = preset
			first_user.Updates(&user)

			return c.JSON(&user)
		})

		app.Put("/api/users", func(c *fiber.Ctx) error {
			user := models.User{}

			if err := c.BodyParser(&user); err != nil {
				return err
			}

			if err := user.UpdateInterval(user.Interval); err != nil {
				return err
			}

			dbUser := models.User{}
			model := database.Database.DB.First(&dbUser, user.ID).Updates(&user)

			if dbUser.IntervalEnabled != user.IntervalEnabled {
				model.Update("interval_enabled", user.IntervalEnabled)
			}

			return c.JSON(dbUser)
		})
	*/
}
