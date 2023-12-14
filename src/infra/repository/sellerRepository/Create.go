package sellerrepository

import (
	"context"
	"log"

	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
	"github.com/pedro-phd/iservice-backend/src/domain/models"
	"github.com/pedro-phd/iservice-backend/src/infra/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sr *SellerRepository) Create(seller models.Seller) (string, *rest_err.RestErr) {

	journey := "SellerRepository CREATE"
	logger.Info("Init GetAll seller", journey)

	db, errConection := db.NewMongoDBConnection(context.Background())
	logger.Info("Database Connected", journey)

	if errConection != nil {
		logger.Error("Error to connect Collection", journey, errConection)
		log.Panic("Error to connect Collection")
	}

	collection := db.Collection("seller")
	logger.Info("Collection Connected", journey)

	sellerToMongo := dto.NewSellerMongoDTO(seller.GetName(), seller.GetEmail(), seller.GetPassword(), seller.GetProducts())

	result, err := collection.InsertOne(context.Background(), sellerToMongo)

	if err != nil {
		return "", rest_err.NewBadRequestError(err.Error())
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil

}
