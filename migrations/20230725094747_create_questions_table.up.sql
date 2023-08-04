CREATE TABLE questions (
	id BIGSERIAL NOT NULL PRIMARY KEY,
    preset_id bigint NOT NULL REFERENCES presets (id),
	title varchar NOT NULL
);