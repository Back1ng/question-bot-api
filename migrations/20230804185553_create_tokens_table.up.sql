ALTER TABLE users ADD CONSTRAINT users_unique_chat_id UNIQUE (chat_id);

CREATE TABLE tokens (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    auth_date bigint NOT NULL,
    first_name varchar(255) NOT NULL,
    hash varchar(255) NOT NULL,
    user_id bigint NOT NULL,
    username varchar(255) NOT NULL
)