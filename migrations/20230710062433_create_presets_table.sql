-- +goose Up
CREATE TABLE presets (
	id bigint NOT NULL PRIMARY KEY,
	title varchar NOT NULL
);

-- +goose Down
DROP TABLE presets;