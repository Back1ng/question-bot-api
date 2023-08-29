package entity

type User struct {
	ID              int64  `json:"id"`
	ChatId          int64  `json:"chat_id"`
	PresetId        int64  `json:"preset_id"`
	Nickname        string `json:"nickname"`
	Interval        int    `json:"interval"`
	IntervalEnabled bool   `json:"interval_enabled"`
}
