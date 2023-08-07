package tgauth_test

import (
	"gitlab.com/back1ng1/question-bot-api/internal/services/tgauth"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	tgbot_token = "6322394098:AAE0gsPnHEv6xHHs4xkf0WIW9JRTLqEY0EQ"
)

func TestAuthIsValid(t *testing.T) {
	os.Setenv("TGBOT_TOKEN", tgbot_token)

	auth := tgauth.Auth{
		AuthDate:  1691177636,
		FirstName: "Кирилл",
		Hash:      "88cf631a05adf9862b56367e84298529a2854cf8c41aa4e40b15bd817cd7d3ca",
		Id:        1258947140,
		Username:  "Ark21bit",
	}

	assert.True(t, auth.IsValid())
}

func TestAuthNotValid(t *testing.T) {
	os.Setenv("TGBOT_TOKEN", tgbot_token)

	// without last char in hash
	auth := tgauth.Auth{
		AuthDate:  1691177636,
		FirstName: "Кирилл",
		Hash:      "88cf631a05adf9862b56367e84298529a2854cf8c41aa4e40b15bd817cd7d3c",
		Id:        1258947140,
		Username:  "Ark21bit",
	}

	assert.False(t, auth.IsValid())
}
