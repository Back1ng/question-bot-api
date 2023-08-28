package question_repository_v1

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	irepository "gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type repository struct {
	db *pgx.Conn
	sb squirrel.StatementBuilderType
}

func NewRepository(db *pgx.Conn, sb squirrel.StatementBuilderType) irepository.QuestionRepository {
	return &repository{db: db, sb: sb}
}

func (r *repository) GetByPreset(presetId int) ([]*entity.Question, error) {
	// todo implement
	return nil, nil
}

func (r *repository) Create(in entity.Question) (*entity.Question, error) {
	// todo implement
	return nil, nil
}

func (r *repository) Update(in entity.Question) (*entity.Question, error) {
	// todo implement
	return nil, nil
}

func (r *repository) Delete(id int) error {
	// todo implement
	return nil
}
