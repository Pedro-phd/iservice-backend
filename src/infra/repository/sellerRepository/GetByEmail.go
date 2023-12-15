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

	temp := &dto.SellerMongoDTO{}

	err := collection.FindOne(context.Background(), bson.D{{"email", email}}).Decode(temp)

	if err != nil {
		return dto.SellerResponseDTO{}, rest_err.NewBadRequestError(err.Error())
	}

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

	return *response, nil
}
