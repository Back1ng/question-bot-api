package repository

import (
	"context"
	"log"

	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/postgres"
)

type UserRepository struct {
	postgres.PgConfig
}

func NewUserRepository(pg postgres.PgConfig) *UserRepository {
	return &UserRepository{pg}
}

func (r UserRepository) FindUserByInterval(i int) ([]entity.User, error) {
	var users []entity.User
	sql, args, err := r.
		Select("id", "chat_id", "nickname", "interval", "interval_enabled").
		From("users").
		Where("interval = ?", i).
		Where("interval = ?", true).
		ToSql()

	if err != nil {
		return users, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		user := entity.User{}
		rows.Scan(&user.ID, &user.ChatId, &user.Nickname, &user.Interval, &user.IntervalEnabled)
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func (r UserRepository) FindUserByChatId(chatId int) (entity.User, error) {
	var user entity.User

	sql, args, err := r.Select("id", "chat_id", "nickname", "interval", "interval_enabled").
		From("users").
		Where("chat_id = ?", chatId).
		ToSql()

	if err != nil {
		return user, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
		args...,
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
	toInsert := map[string]interface{}{
		"chat_id":   u.ChatId,
		"nickname":  u.Nickname,
		"preset_id": u.PresetId,
	}

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
		return u, CreateUserError
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
		return user, UpdateUserError
	}

	return u, nil
}
