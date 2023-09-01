package crud_presets

import (
	"gitlab.com/back1ng1/question-bot-api/app/repository"
	"gitlab.com/back1ng1/question-bot-api/entity"
	"gitlab.com/back1ng1/question-bot-api/pkg/logger"
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
		logger.Log.Errorf(
			"app.usecase.crud_presets.implement.GetAll() - uc.presetRepo.GetAll(): %v",
			err,
		)

		return nil, err
	}

	return out, nil
}

func (uc *usecase) Create(in entity.Preset) (*entity.Preset, error) {
	out, err := uc.presetRepo.Create(in)

	if err != nil {
		logger.Log.Errorf(
			"app.usecase.crud_presets.implement.Create() - uc.presetRepo.Create(in): %v. Preset: %#+v",
			err,
			in,
		)

		return nil, err
	}

	return out, nil
}

func (uc *usecase) Update(in entity.Preset) (*entity.Preset, error) {
	out, err := uc.presetRepo.Update(in)

	if err != nil {
		logger.Log.Errorf(
			"app.usecase.crud_presets.implement.Update() - uc.presetRepo.Update(in): %v. Preset: %#+v",
			err,
			in,
		)

		return nil, err
	}

	return out, nil
}

func (uc *usecase) Delete(id int64) error {
	err := uc.presetRepo.Delete(id)

	if err != nil {
		logger.Log.Errorf(
			"app.usecase.crud_presets.implement.Delete() - uc.presetRepo.Delete(id): %v. id: %d",
			err,
			id,
		)
	}

	return err
}
