package sellerRepository

import (
	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
)

func (sr *SellerRepository) FindAll() ([]dto.SellerResponse, error) {

	journey := "FindAll - Seller"
	logger.Info("Init to FindAll", journey)

	var results []dto.SellerResponse

	err := sr.db.Select(&results, "SELECT * FROM sellers")

	if err != nil {
		logger.Error("Error to find all sellers", journey, err)
		return nil, err
	}
	return results, nil
}
