package user_repository

import (
	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
)

func FindByInterval(i int) []models.User {
	users := []models.User{}

	database.Database.DB.Find(
		&users,
		models.User{
			Interval:        i,
			IntervalEnabled: true,
		},
	)

	return users
}
