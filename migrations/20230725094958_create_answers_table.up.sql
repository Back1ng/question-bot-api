CREATE TABLE answers (
	id bigint NOT NULL PRIMARY KEY,
    question_id bigint NOT NULL,
	title varchar NOT NULL,
    is_correct bool DEFAULT 'false'
);