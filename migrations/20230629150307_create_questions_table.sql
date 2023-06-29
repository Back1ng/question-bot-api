-- +goose Up
CREATE TABLE questions (
	id int NOT NULL PRIMARY KEY
);

-- +goose Down
DROP TABLE questions;