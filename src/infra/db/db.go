package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
)

func NewDBConnection() *sqlx.DB {
	journey := "NewDBConnection"
	logger.Info("Init to create connection with database", journey)

	dsn := os.Getenv("DB_URL")

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		logger.Error("Error to open connect with database", journey, err)
	}

	errPing := db.Ping()

	if errPing != nil {
		logger.Error("Error to establishing connection if database", journey, errPing)
		panic("Error to establishing connection if database")
	}

	return db
}
