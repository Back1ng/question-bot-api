package repository

import (
	"context"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"log"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	*pgx.Conn
}

func NewUserRepository(conn *pgx.Conn) *UserRepository {
	return &UserRepository{conn}
}

func (r UserRepository) UserFindByInterval(i int) []entity.User {
	rows, err := r.Query(
		context.Background(),
		`SELECT id, chat_id, nickname, interval, interval_enabled 
		FROM users 
		WHERE interval = @interval AND interval_enabled = true`,
		pgx.NamedArgs{
			"interval": i,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []entity.User

	for rows.Next() {
		user := entity.User{}
		rows.Scan(&user.ID, &user.ChatId, &user.Nickname, &user.Interval, &user.IntervalEnabled)
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users
}

func (r UserRepository) CreateUser(u entity.User) (entity.User, error) {
	return u, nil
}

func (r UserRepository) FindUserByChatId(chatId int) (entity.User, error) {
	var user entity.User

	rows, err := r.Query(
		context.Background(),
		`SELECT id, chat_id, nickname, interval, interval_enabled
			FROM users
			WHERE chat_id = @chat_id`,
		pgx.NamedArgs{
			"chat_id": chatId,
		},
	)
	if err != nil {
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.ID, &user.ChatId, &user.Nickname, &user.Interval, &user.IntervalEnabled)
		break
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return user, err
	}

	return user, nil
}
