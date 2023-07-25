CREATE TABLE answers (
	id BIGSERIAL NOT NULL PRIMARY KEY,
    question_id bigint NOT NULL REFERENCES questions (id),
	title varchar NOT NULL,
    is_correct bool DEFAULT 'false'
);