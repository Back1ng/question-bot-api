package crud_presets

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
)

type usecase struct {
	presetRepo repository.PresetRepository
}

func NewUseCase(presetRepo repository.PresetRepository) UseCase {
	return &usecase{presetRepo: presetRepo}
}

func (uc *usecase) GetAll() ([]*entity.Preset, error) {
	out, err := uc.presetRepo.GetAll()

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Create(in entity.Preset) (*entity.Preset, error) {
	out, err := uc.presetRepo.Create(in)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Update(in entity.Preset) (*entity.Preset, error) {
	out, err := uc.presetRepo.Update(in)

	if err != nil {
		return nil, err
	}

	return out, nil
}

func (uc *usecase) Delete(id int64) error {
	return uc.presetRepo.Delete(id)
}
