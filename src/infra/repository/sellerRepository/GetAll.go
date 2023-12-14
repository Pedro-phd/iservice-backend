package sellerrepository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
	"github.com/pedro-phd/iservice-backend/src/infra/db"
	"go.mongodb.org/mongo-driver/bson"
)

func (sr *SellerRepository) GetAll() ([]dto.SellerResponseDTO, *rest_err.RestErr) {

	journey := "SellerRepository GETALL"
	logger.Info("Init GetAll seller", journey)

	db, errConection := db.NewMongoDBConnection(context.Background())
	logger.Info("Database Connected", journey)

	if errConection != nil {
		logger.Error("Error to connect Collection", journey, errConection)
		log.Panic("Error to connect Collection")
	}

	collection := db.Collection("seller")
	logger.Info("Collection Connected", journey)

	var results []dto.SellerMongoDTO
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		logger.Error("Error to colect data", journey, err)
		log.Panic("Error to colect data")
		return nil, rest_err.NewBadRequestError(err.Error())
	}

	logger.Info("Data Collected", journey)

	if errCursor := cursor.All(context.Background(), &results); errCursor != nil {
		logger.Error("Error to iterate data", journey, errCursor)
		return nil, rest_err.NewBadRequestError(errCursor.Error())
	}

	var response []dto.SellerResponseDTO
	for _, seller := range results {
		cursor.Decode(&seller)
		out, err := json.Marshal(seller)
		if err != nil {
			logger.Error("Error to marshal", journey, err)
			return nil, rest_err.NewBadRequestError(err.Error())
		}
		fmt.Printf("%s\n", out)
		var res dto.SellerResponseDTO
		json.Unmarshal(out, &res)
		response = append(response, res)
	}

	return response, nil
}
