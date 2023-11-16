package user_repository_v1

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"

	"github.com/Masterminds/squirrel"
	irepository "gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
)

type repository struct {
	db *pgxpool.Pool
	sb squirrel.StatementBuilderType
}

func NewRepository(db *pgxpool.Pool, sb squirrel.StatementBuilderType) irepository.UserRepository {
	return &repository{
		db: db,
		sb: sb,
	}
}

func (r *repository) GetByInterval(interval int) ([]int64, error) {
	var users []int64
	sql, args, err := r.sb.
		Select("chat_id").
		From("users").
		Where("interval = ?", interval).
		Where("interval_enabled = ?", true).
		ToSql()
	if err != nil {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.GetByInterval() - r.sb.Select(): %v",
			err,
		)
		return users, err
	}

	rows, err := r.db.Query(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.GetByInterval() - r.db.Query(): %v",
			err,
		)
		return users, err
	}

	defer rows.Close()

	for rows.Next() {
		var chatId int64
		rows.Scan(&chatId)
		users = append(users, chatId)
	}

	if err := rows.Err(); err != nil {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.GetByInterval() - rows.Err(): %v",
			err,
		)
		return users, err
	}

	return users, nil
}

func (r *repository) GetByChatId(chatId int) (*entity.User, error) {
	sql, args, err := r.sb.Select("id", "chat_id", "preset_id", "nickname", "interval", "interval_enabled").
		From("users").
		Where("chat_id = ?", chatId).
		ToSql()
	if err != nil {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.GetByChatId() - r.sb.Select(): %v. ChatId: %d",
			err,
			chatId,
		)
		return nil, err
	}

	rows, err := r.db.Query(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.GetByChatId() - r.db.Query(): %v. ChatId: %d",
			err,
			chatId,
		)
		return nil, err
	}

	defer rows.Close()

	var user entity.User
	for rows.Next() {
		rows.Scan(&user.ID, &user.ChatId, &user.PresetId, &user.Nickname, &user.Interval, &user.IntervalEnabled)
		break
	}

	if err := rows.Err(); err != nil {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.GetByChatId() - rows.Err(): %v. ChatId: %d",
			err,
			chatId,
		)
		return nil, err
	}

	return &user, nil
}

func (r *repository) Create(in entity.User) (*entity.User, error) {
	toInsert := map[string]interface{}{
		"chat_id":   in.ChatId,
		"nickname":  in.Nickname,
		"preset_id": in.PresetId,
	}

	if in.Interval != 0 {
		toInsert["interval"] = in.Interval
	}

	sql, args, err := r.sb.Insert("users").
		Suffix("RETURNING id").
		SetMap(toInsert).
		ToSql()
	if err != nil {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.Create() - r.sb.Insert(): %v. User: %#+v",
			err,
			in,
		)
		return nil, err
	}

	row := r.db.QueryRow(
		context.Background(),
		sql,
		args...,
	)

	if err := row.Scan(&in.ID); err != nil {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.Create() - row.Scan(): %v. User: %#+v",
			err,
			in,
		)
		return nil, err
	}

	return &in, nil
}

func (r *repository) Update(in entity.User) (*entity.User, error) {
	query := r.sb.Update("users").
		Where("chat_id = ?", in.ChatId)

	i := reflect.ValueOf(in.IntervalEnabled)
	if !i.IsZero() {
		query.Set("interval_enabled", in.IntervalEnabled)
	}

	if in.Interval > 0 && in.Interval < 25 {
		query = query.Set("interval", in.Interval)
	}

	if in.PresetId > 0 {
		query = query.Set("preset_id", in.PresetId)
	}

	sql, args, err := query.ToSql()
	if err != nil {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.Update() - query.ToSql(): %v. User: %#+v",
			err,
			in,
		)
		return nil, err
	}

	commandTag, err := r.db.Exec(
		context.Background(),
		sql,
		args...,
	)
	if err != nil {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.Update() - r.db.Exec(): %v. User: %#+v",
			err,
			in,
		)
		return nil, err
	}

	if commandTag.RowsAffected() != 1 {
		logger.Log.Errorf(
			"repository.user_repository_v1.repository.Update() - r.db.Exec(): %v. User: %#+v",
			"update user: no affected rows",
			in,
		)
		return nil, errors.New("update user: no affected rows")
	}

	return &in, nil
}
