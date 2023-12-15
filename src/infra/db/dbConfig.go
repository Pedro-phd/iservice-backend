package db

import (
	"context"
	"log"

	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDbConfig() {

	journey := "initDbConfig"

	logger.Info("Init to configuration of db", journey)
	db, err := NewMongoDBConnection(context.Background())

	if err != nil {
		logger.Error("Error to connect with db", journey, err)
		log.Fatal("Error to connect with db")
	}

	indexName := mongo.IndexModel{
		Keys:    bson.D{{Key: "name", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	indexEmail := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	// indexDocument := mongo.IndexModel{
	// 	Keys:    bson.D{{Key: "document", Value: 1}},
	// 	Options: options.Index().SetUnique(true),
	// }

	indexs := []mongo.IndexModel{indexName, indexEmail}

	collection := db.Collection("seller")

	logger.Info("Init insert indexes in db", journey)
	_, errIndexs := collection.Indexes().CreateMany(context.Background(), indexs)

	if errIndexs != nil {
		logger.Error("Error to create indexes in db", journey, errIndexs)
		log.Fatal("Error to create indexes in db")
	}

	defer db.Client().Disconnect(context.Background())

	logger.Info("Finish db configuration", journey)
}
