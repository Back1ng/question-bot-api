# question-bot

Миграции запускаются по команде из контейнера:
migrate -database ${POSTGRESQL_URL} -path migrations/ up