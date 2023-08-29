package crud_presets

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type usecase struct {
	preset_repo repository.PresetRepository
}

func NewUseCase(preset_repo repository.PresetRepository) UseCase {
	return &usecase{preset_repo: preset_repo}
}

func (uc *usecase) GetAll() ([]*entity.Preset, error) {
	out, err := uc.preset_repo.GetAll()

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Create(in entity.Preset) (*entity.Preset, error) {
	out, err := uc.preset_repo.Create(in)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Update(in entity.Preset) (*entity.Preset, error) {
	out, err := uc.preset_repo.Update(in)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Delete(id int64) error {
	return uc.preset_repo.Delete(id)
}
