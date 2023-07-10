-- +goose Up
CREATE TABLE users (
	id bigint NOT NULL PRIMARY KEY,
	chat_id bigint NOT NULL,
	name varchar NOT NULL,
	preset_id bigint
);

-- +goose Down
DROP TABLE users;