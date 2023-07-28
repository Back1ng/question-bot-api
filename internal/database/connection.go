package database

import "github.com/jackc/pgx/v5"

type DbInstance struct {
	*pgx.Conn
}

var Database DbInstance
