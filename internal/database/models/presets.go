package models

type Preset struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Questions []Question
}
