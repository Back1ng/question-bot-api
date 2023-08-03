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

func (r UserRepository) FindUserByInterval(i int) []entity.User {
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

func (r UserRepository) CreateUser(u entity.User) (entity.User, error) {
	toInsert := map[string]interface{}{}

	toInsert["chat_id"] = u.ChatId
	toInsert["nickname"] = u.Nickname
	toInsert["preset_id"] = u.PresetId

	if u.Interval != 0 {
		toInsert["interval"] = u.Interval
	}

	sql, args, err := r.Insert("users").SetMap(toInsert).ToSql()
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

func (r UserRepository) UpdateUser(u entity.User) (entity.User, error) {
	var user entity.User

	query := r.Update("users").
		Where("chat_id = ?", u.ChatId).
		Set("interval_enabled", u.IntervalEnabled)

	if u.Interval > 0 && u.Interval < 25 {
		query = query.Set("interval", u.Interval)
	}

	sql, args, err := query.ToSql()

	if err != nil {
		return user, err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		return user, err
	}

	if commandTag.RowsAffected() != 1 {
		return user, errors.New("cannot update user")
	}

	return u, nil
}
