-- +goose Up
CREATE TABLE questions (
	id bigint NOT NULL PRIMARY KEY,
	title varchar NOT NULL
);

-- +goose Down
DROP TABLE questions;