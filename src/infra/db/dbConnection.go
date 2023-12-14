package db

import (
	"context"
	"os"

	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {

	journey := "InitDBConnection"

	logger.Info("Init Database Connection", journey)

	dburl := os.Getenv("DB_URL")
	dbname := os.Getenv("DB_NAME")

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburl))

	if err != nil {
		logger.Error("Database Connection failed", journey, err)
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Database ping failed", journey, err)
		return nil, err
	}

	db := client.Database(dbname)

	logger.Info("Database Connection successfully", journey)
	return db, nil
}
