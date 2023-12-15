package sellerrepository

import (
	"context"
	"log"

	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
	"github.com/pedro-phd/iservice-backend/src/infra/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sc *SellerRepository) GetById(id string) (dto.SellerResponseDTO, *rest_err.RestErr) {
	journey := "SellerRepository-getById"
	logger.Info("Init GetById seller", journey)

	db, errConection := db.NewMongoDBConnection(context.Background())
	logger.Info("Database Connected", journey)

	if errConection != nil {
		logger.Error("Error to connect Collection", journey, errConection)
		log.Panic("Error to connect Collection")
	}

	collection := db.Collection("seller")
	logger.Info("Collection Connected", journey)

	temp := &dto.SellerMongoDTO{}

	idPrimitive, _ := primitive.ObjectIDFromHex(id)
	err := collection.FindOne(context.Background(), bson.D{{"_id", idPrimitive}}).Decode(temp)

	var prodsDTO []dto.ProductResponseDTO

	for _, prod := range temp.Products {
		temp := dto.ProductResponseDTO{
			Name:        prod.Name,
			Price:       prod.Price,
			Description: prod.Description,
			Thumb:       prod.Thumb,
			CreatedAt:   prod.CreatedAt,
			UpdatedAt:   prod.UpdatedAt,
		}

		prodsDTO = append(prodsDTO, temp)
	}

	response := &dto.SellerResponseDTO{
		ID:        temp.ID,
		Name:      temp.Name,
		Email:     temp.Password,
		Products:  prodsDTO,
		CreatedAt: temp.CreatedAt,
		UpdatedAt: temp.UpdatedAt,
	}

	if err != nil {
		return dto.SellerResponseDTO{}, rest_err.NewBadRequestError(err.Error())
	}

	return *response, nil
}
