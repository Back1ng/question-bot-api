package repository

import "errors"

var (
	// answers
	UpdateAnswerEmptyTitle     = errors.New("update answer: title is null in update answer")
	UpdateAnswerError          = errors.New("update answer: cannot update answers")
	DeleteAnswerIdNotPresented = errors.New("delete answer: id is not presented")
	DeleteAnswerError          = errors.New("delete answer: cannot delete answer")

	// auth
	AuthDataIsOutdated   = errors.New("auth: data is outdated")
	AuthFailedCheck      = errors.New("auth: failed check auth")
	AuthCannotStoreToken = errors.New("auth: cannot store token")

	// presets
	UpdatePresetsEmptyTitle = errors.New("update preset: title is null")
	UpdatePresetsError      = errors.New("update preset: cannot update presets")
	DeletePresetsError      = errors.New("delete preset: cannot delete preset")

	// questions
	UpdateQuestionEmptyTitle = errors.New("update question: title is null")
	UpdateQuestionError      = errors.New("update question: cannot update question")
	DeleteQuestionError      = errors.New("delete question: cannot delete question")

	// users
	CreateUserError = errors.New("create user: user cannot be created")
	UpdateUserError = errors.New("update user: user cannot be updated")
)
