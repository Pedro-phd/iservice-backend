package sellerRepository

import (
	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
	"github.com/pedro-phd/iservice-backend/src/domain/models"
)

func (sr *SellerRepository) FindById(sellerId string) (dto.SellerResponse, error) {

	journey := "FindById - Seller"
	logger.Info("Init to find by id seller", journey)

	var result []dto.SellerResponse
	var categoriesResult []models.SellersCategories

	err := sr.db.Select(&result, "SELECT * FROM sellers WHERE id = $1", sellerId)

	if err != nil {
		logger.Error("Error to find by id seller", journey, err)
		return dto.SellerResponse{}, err
	}

	errCategories := sr.db.Select(
		&categoriesResult,
		"SELECT * FROM sellers_categories WHERE seller_id = $1",
		sellerId,
	)

	if errCategories != nil {
		logger.Error("Error to find categories by id seller", journey, errCategories)
		return dto.SellerResponse{}, errCategories
	}

	if len(result) == 0 {
		return dto.SellerResponse{}, rest_err.NewNotFoundError("Not found seller")
	}

	result[0].Categories = categoriesResult

	return result[0], nil
}
