package database

import "github.com/jackc/pgx/v5"

type DbInstance struct {
	DB *pgx.Conn
}

var Database DbInstance
