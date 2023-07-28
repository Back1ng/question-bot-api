package database

import "gitlab.com/back1ng1/question-bot/internal/database/entity"

type AnswerRepository interface {
	FindAnswersInQuestion(questionId int) ([]entity.Answer, error)
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
}

type Repositories interface {
	AnswerRepository
	PresetRepository
	QuestionRepository
	UserRepository
}
