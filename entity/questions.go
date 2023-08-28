package entity

type Question struct {
	ID       int64  `json:"id"`
	PresetId int64  `json:"preset_id"`
	Title    string `json:"title"`
}
