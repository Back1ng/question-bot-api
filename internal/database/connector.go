package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbInstance struct {
	DB *gorm.DB
}

var Database DbInstance

func SetupConnection() {
	dsn := fmt.Sprintf("postgresql://%v:%v@db:%v/%v", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("cannot connect to database")
	}

	Database = DbInstance{db}
}
