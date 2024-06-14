# question-bot

## This is clone Golang REST API from private gitlab
Миграции запускаются по команде из контейнера:
migrate -database ${POSTGRESQL_URL} -path migrations/ up