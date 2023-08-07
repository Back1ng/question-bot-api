package database

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/internal/database/repository"
	"gitlab.com/back1ng1/question-bot-api/internal/services/tgauth"
	"gitlab.com/back1ng1/question-bot-api/pkg/postgres"
)

type AnswerRepository interface {
	FindAnswersInQuestion(questionId int) ([]entity.Answer, error)
	StoreAnswer(answer entity.Answer) (entity.Answer, error)
	UpdateAnswer(answer entity.Answer) error
	DeleteAnswer(answer entity.Answer) error
}

type PresetRepository interface {
	FindPresets() ([]entity.Preset, error)
	StorePreset(p entity.Preset) (entity.Preset, error)
	UpdatePreset(id int, p entity.Preset) error
	DeletePreset(id int) error
}

type QuestionRepository interface {
	FindQuestionsInPreset(presetId int) ([]entity.Question, error)
	StoreQuestion(q entity.Question) error
	UpdateQuestionTitle(id int, q entity.Question) error
	DeleteQuestion(id int) error
}

type UserRepository interface {
	FindUserByInterval(i int) ([]entity.User, error)
	FindUserByChatId(chatId int) (entity.User, error)
	CreateUser(u entity.User) (entity.User, error)
	UpdateUser(u entity.User) (entity.User, error)
}

type AuthRepository interface {
	HasToken(hash string) (bool, error)
	GenerateToken(auth tgauth.Auth) (string, error)
}

type Repositories struct {
	AnswerRepository
	PresetRepository
	QuestionRepository
	UserRepository
	AuthRepository
}

func GetRepositories(conn *pgx.Conn, sb squirrel.StatementBuilderType) Repositories {
	pg := postgres.PgConfig{Conn: conn, StatementBuilderType: sb}

	return Repositories{
		AnswerRepository:   repository.NewAnswerRepository(pg),
		PresetRepository:   repository.NewPresetRepository(pg),
		QuestionRepository: repository.NewQuestionRepository(pg),
		UserRepository:     repository.NewUserRepository(pg),
		AuthRepository:     repository.NewAuthRepository(pg),
	}
}
