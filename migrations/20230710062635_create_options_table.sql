-- +goose Up
CREATE TABLE options (
	id bigint NOT NULL PRIMARY KEY,
	question_id bigint NOT NULL,
	title varchar NOT NULL,
	is_correct boolean NOT NULL default false
);

-- +goose Down
DROP TABLE options;