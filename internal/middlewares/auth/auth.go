package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.Filter(c) {
			return c.Next()
		}

		headers := c.GetReqHeaders()

		if headers["Authorization"] == "" {
			return errors.New("auth: no valid authorization")
		}

		authString := strings.Split(headers["Authorization"], " ")
		if len(authString) != 2 {
			return errors.New("auth: wrong authentication header. expected: bearer token")
		}

		hasToken, err := config.Repo.HasToken(authString[1])
		if err != nil {
			return err
		}

		if hasToken {
			return c.Next()
		}

		return errors.New("auth: authorization token not valid")
	}
}
