package auth

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
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
			logger.Log.Errorf(
				"internal.middlewares.auth.New() - headers[\"\"] == \"\": %v",
				"No valid authorization",
			)
			return ErrorNoValidAuth
		}

		authString := strings.Split(headers["Authorization"], " ")
		if len(authString) != 2 {
			logger.Log.Errorf(
				"internal.middlewares.auth.New() - len(authString) != 2: %v. Given: %v",
				"Wrong authentication header. Expected: Bearer Token", headers["Authorization"],
			)
			return ErrorWrongHeader
		}

		token, err := config.Repo.Get(authString[1])
		if err != nil {
			logger.Log.Error(err)
			return err
		}

		if token.Hash != "" {
			return c.Next()
		}

		logger.Log.Errorf(
			"internal.middlewares.auth.New() - eof: %v",
			"authorization token not valid",
		)
		return ErrorTokenNotValid
	}
}
