CREATE TABLE users (
    id bigint NOT NULL PRIMARY KEY,
    chat_id bigint NOT NULL,
    preset_id bigint NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    interval int,
    interval_enabled bool DEFAULT false
)