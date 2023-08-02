package repository

import (
	"context"
	"errors"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/postgres"
	"log"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	postgres.PgConfig
}

func NewUserRepository(pg postgres.PgConfig) *UserRepository {
	return &UserRepository{pg}
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
	query := r.Insert("users").
		Columns("chat_id", "nickname", "preset_id").
		Values(u.ChatId, u.Nickname, u.PresetId)

	if u.Interval != 0 {
		query.Columns("interval").Values(u.Interval)
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return u, err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		return u, err
	}

	if commandTag.RowsAffected() != 1 {
		return u, errors.New("user cannot be created")
	}

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
