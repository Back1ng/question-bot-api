package crud_presets

import "gitlab.com/back1ng1/question-bot-api/entity"

type UseCase interface {
	GetAll() ([]*entity.Preset, error)
	Create(in entity.Preset) (*entity.Preset, error)
	Update(in entity.Preset) (*entity.Preset, error)
	Delete(id int64) error
}
