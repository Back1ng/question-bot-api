package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var (
	ErrorNoValidAuth   = errors.New("auth: no valid authorization")
	ErrorTokenNotValid = errors.New("auth: authorization token not valid")
	ErrorWrongHeader   = errors.New("auth: wrong authentication header. expected: bearer token")
)

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if config.Filter(c) {
			return c.Next()
		}

		headers := c.GetReqHeaders()

		if headers["Authorization"] == "" {
			return ErrorNoValidAuth
		}

		authString := strings.Split(headers["Authorization"], " ")
		if len(authString) != 2 {
			return ErrorWrongHeader
		}

		token, err := config.Repo.Get(authString[1])
		if err != nil {
			return err
		}

		if token.Hash != "" {
			return c.Next()
		}

		return ErrorTokenNotValid
	}
}
