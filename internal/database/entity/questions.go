package entity

type Question struct {
	ID       int64 `json:"id"`
	PresetId int64 `json:"preset_id"`
	Preset   Preset
	Title    string   `json:"title" validate:"required"`
	Answers  []Answer `json:"answers" validate:"required"`
}
