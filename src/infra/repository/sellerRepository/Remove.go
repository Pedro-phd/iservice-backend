package sellerrepository

import (
	"context"
	"log"

	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/infra/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sc *SellerRepository) Remove(id string) bool {

	journey := "SellerRepository REMOVE"
	logger.Info("Init Remove seller", journey)

	db, errConection := db.NewMongoDBConnection(context.Background())
	logger.Info("Database Connected", journey)

	if errConection != nil {
		logger.Error("Error to connect Collection", journey, errConection)
		log.Panic("Error to connect Collection")
	}

	collection := db.Collection("seller")
	logger.Info("Collection Connected", journey)

	idPrimitive, _ := primitive.ObjectIDFromHex(id)
	result, err := collection.DeleteOne(context.Background(), bson.D{{"_id", idPrimitive}})

	if err != nil {
		logger.Error("Error to delete", journey, err)
	}

	if result.DeletedCount <= 0 {
		return false
	}

	return true
}
