CREATE TABLE users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    preset_id bigint NOT NULL REFERENCES presets (id),
    chat_id bigint NOT NULL,
    nickname VARCHAR(255) NOT NULL,
    interval int DEFAULT 3 NOT NULL,
    interval_enabled bool DEFAULT false
)