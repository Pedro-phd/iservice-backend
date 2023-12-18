package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
)

func NewDBConnection() *sqlx.DB {
	journey := "NewDBConnection"
	dsn := os.Getenv("DB_URL")

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		logger.Error("Error to open connect with database", journey, err)
	}

	return db
}
