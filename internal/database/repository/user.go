package repository

import (
	"context"
	"log"

	"gitlab.com/back1ng1/question-bot/internal/database"
	"gitlab.com/back1ng1/question-bot/internal/database/models"
)

func UserFindByInterval(i int) []models.User {
	rows, err := database.Database.DB.Query(
		context.Background(),
		"SELECT id, chat_id, nickname, interval, interval_enabled FROM users WHERE interval=$1 AND interval_enabled=$2",
		i,
		true,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []models.User{}

	for rows.Next() {
		user := models.User{}

		rows.Scan(&user.ID, &user.ChatId, &user.Nickname, &user.Interval, &user.IntervalEnabled)

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}
