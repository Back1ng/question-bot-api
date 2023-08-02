package database

import (
	"github.com/jackc/pgx/v5"
	"gitlab.com/back1ng1/question-bot-api/internal/database/entity"
	"gitlab.com/back1ng1/question-bot-api/internal/database/repository"
)

type AnswerRepository interface {
	FindAnswersInQuestion(questionId int) ([]entity.Answer, error)
	StoreAnswer(answer entity.Answer) (entity.Answer, error)
	UpdateAnswer(answer entity.Answer) error
	DeleteAnswer(answer entity.Answer) error
}

type PresetRepository interface {
	FindPresets() ([]entity.Preset, error)
	StorePreset(p entity.Preset) error
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
	UserFindByInterval(i int) []entity.User
	CreateUser(u entity.User) (entity.User, error)
	FindUserByChatId(chatId int) (entity.User, error)
}

type Repositories struct {
	AnswerRepository
	PresetRepository
	QuestionRepository
	UserRepository
}

func GetRepositories(conn *pgx.Conn) Repositories {
	return Repositories{
		AnswerRepository:   repository.NewAnswerRepository(conn),
		PresetRepository:   repository.NewPresetRepository(conn),
		QuestionRepository: repository.NewQuestionRepository(conn),
		UserRepository:     repository.NewUserRepository(conn),
	}
}
