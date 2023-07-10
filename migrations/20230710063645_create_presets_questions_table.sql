-- +goose Up
CREATE TABLE presets_questions (
	preset_id bigint NOT NULL,
	question_id bigint NOT NULL
);

-- +goose Down
DROP TABLE presets_questions;