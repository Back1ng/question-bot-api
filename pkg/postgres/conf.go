package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

type PgConfig struct {
	*pgx.Conn
	squirrel.StatementBuilderType
}
