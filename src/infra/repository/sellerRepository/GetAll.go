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

	// var results []dto.SellerMongoDTO
	cursor, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		logger.Error("Error to colect data", journey, err)
		log.Panic("Error to colect data")
		return nil, rest_err.NewBadRequestError(err.Error())
	}

	logger.Info("Data Collected", journey)

	var sellerData []dto.SellerMongoDTO

	if err := cursor.All(context.Background(), &sellerData); err != nil {
		logger.Error("Error to iterate data", journey, err)
		return nil, rest_err.NewBadRequestError(err.Error())
	}

	var result []dto.SellerResponseDTO

	for _, seller := range sellerData {

		prodsResponse := make([]dto.ProductResponseDTO, 0)

		for _, prod := range seller.Products {
			temp := dto.ProductResponseDTO{
				Name:        prod.Name,
				Price:       prod.Price,
				Description: prod.Description,
				Thumb:       prod.Thumb,
				CreatedAt:   prod.CreatedAt,
				UpdatedAt:   prod.UpdatedAt,
			}
			prodsResponse = append(prodsResponse, temp)
		}

		temp := dto.SellerResponseDTO{
			ID:        seller.ID,
			Name:      seller.Name,
			Email:     seller.Email,
			Products:  prodsResponse,
			CreatedAt: seller.CreatedAt,
			UpdatedAt: seller.UpdatedAt,
		}
		result = append(result, temp)
	}

	logger.Info("Data Converted", journey)

	defer db.Client().Disconnect(context.Background())
	return result, nil
}
