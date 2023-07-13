package presets

import "github.com/gofiber/fiber/v2"

func Routes(app *fiber.App) {
	app.Get("/api/presets", func(c *fiber.Ctx) error {
		// send presets
		return c.SendString("test")
	})

	app.Post("/api/presets", func(c *fiber.Ctx) error {
		// store preset
		payload := struct {
			Name string `json:"name"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		return c.JSON(payload)
	})
}
