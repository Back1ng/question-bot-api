package tgauth_test

import (
	"os"
	"testing"

	"gitlab.com/back1ng1/question-bot-api/pkg/tgauth"

	"github.com/stretchr/testify/assert"
)

const (
	tgbotToken = "6322394098:AAE0gsPnHEv6xHHs4xkf0WIW9JRTLqEY0EQ"
)

func TestAuthIsValid(t *testing.T) {
	os.Setenv("TGBOT_TOKEN", tgbotToken)

	tests := []struct {
		name   string
		give   tgauth.Auth
		wanted bool
	}{
		{
			name: "assert auth is passed",
			give: tgauth.Auth{
				AuthDate:  1691177636,
				FirstName: "Кирилл",
				Hash:      "88cf631a05adf9862b56367e84298529a2854cf8c41aa4e40b15bd817cd7d3ca",
				Id:        1258947140,
				Username:  "Ark21bit",
			},
			wanted: true,
		},
		{
			name: "assert auth is not passed",
			give: tgauth.Auth{
				AuthDate:  1691177636,
				FirstName: "Кирилл",
				Hash:      "88cf631a05adf9862b56367e84298529a2854cf8c41aa4e40b15bd817cd7d3c", // without last char
				Id:        1258947140,
				Username:  "Ark21bit",
			},
			wanted: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.give.IsValid(), tt.wanted)
		})
	}
}
