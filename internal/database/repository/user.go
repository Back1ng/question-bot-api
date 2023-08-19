package repository

import (
	"context"

	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
	"gitlab.com/back1ng1/question-bot-api/pkg/postgres"
)

type UserRepository struct {
	postgres.PgConfig
}

func NewUserRepository(pg postgres.PgConfig) *UserRepository {
	return &UserRepository{pg}
}

func (r UserRepository) FindUsersByInterval(i int) ([]int64, error) {
	var users []int64
	sql, args, err := r.
		Select("chat_id").
		From("users").
		Where("interval = ?", i).
		Where("interval_enabled = ?", true).
		ToSql()

	if err != nil {
		logger.Log.Errorf("UserRepository.FindUsersByInterval - r.Select: %v", err)
		return users, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf("UserRepository.FindUsersByInterval - r.Query: %v", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var chatId int64
		rows.Scan(&chatId)
		users = append(users, chatId)
	}

	if err := rows.Err(); err != nil {
		logger.Log.Errorf("UserRepository.FindUserByInterval - rows.err: %v", err)
		return users, err
	}

	return users, nil
}

func (r UserRepository) FindUserByChatId(chatId int) (entity.User, error) {
	var user entity.User

	sql, args, err := r.Select("id", "chat_id", "preset_id", "nickname", "interval", "interval_enabled").
		From("users").
		Where("chat_id = ?", chatId).
		ToSql()

	if err != nil {
		logger.Log.Errorf("UserRepository.FindUserByChatId - r.Select: %v", err)
		return user, err
	}

	rows, err := r.Query(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf("UserRepository.FindUserByChatId - r.Query: %v", err)
		return user, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&user.ID, &user.ChatId, &user.PresetId, &user.Nickname, &user.Interval, &user.IntervalEnabled)
		break
	}

	if err := rows.Err(); err != nil {
		logger.Log.Errorf("UserRepository.FindUserByChatId - rows.Err: %v", err)
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

	sql, args, err := r.Insert("users").
		Suffix("RETURNING id").
		SetMap(toInsert).
		ToSql()
	if err != nil {
		logger.Log.Errorf("UserRepository.CreateUser - r.Insert: %v", err)
		return u, err
	}

	row := r.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	if err := row.Scan(&u.ID); err != nil {
		logger.Log.Errorf("UserRepository.CreateUser - r.QueryRow: %v", err)
		return u, err
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

	if u.PresetId > 0 {
		query = query.Set("preset_id", u.PresetId)
	}

	sql, args, err := query.ToSql()

	if err != nil {
		logger.Log.Errorf("UserRepository.UpdateUser - query.ToSql: %v", err)
		return user, err
	}

	commandTag, err := r.Exec(
		context.Background(),
		sql,
		args...,
	)

	if err != nil {
		logger.Log.Errorf("UserRepository.UpdateUser - r.Exec: %v", err)
		return user, err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf("UserRepository.UpdateUser - query.ToSql: %v", UpdateUserError)
		return user, UpdateUserError
	}

	return u, nil
}
