package sellerrepository

import (
	"context"
	"log"

	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
	"github.com/pedro-phd/iservice-backend/src/infra/db"
	"go.mongodb.org/mongo-driver/bson"
)

func (sc *SellerRepository) GetByEmail(email string) (dto.SellerResponseDTO, *rest_err.RestErr) {
	journey := "SellerRepository GETBYPARAM"
	logger.Info("Init GetByParam seller", journey)

	db, errConection := db.NewMongoDBConnection(context.Background())
	logger.Info("Database Connected", journey)

	if errConection != nil {
		logger.Error("Error to connect Collection", journey, errConection)
		log.Panic("Error to connect Collection")
	}

	collection := db.Collection("seller")
	logger.Info("Collection Connected", journey)

	response := &dto.SellerResponseDTO{}

	err := collection.FindOne(context.Background(), bson.D{{"email", email}}).Decode(response)

	if err != nil {
		return dto.SellerResponseDTO{}, rest_err.NewBadRequestError(err.Error())
	}

	return *response, nil
}
