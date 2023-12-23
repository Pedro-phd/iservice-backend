package sellerRepository

import (
	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
)

func (sr *SellerRepository) Delete(sellerId int) error {

	journey := "Delete - Seller"
	logger.Info("Init to delete seller", journey)

	_, err := sr.db.Exec("DELETE FROM sellers WHERE id = $1", sellerId)
	if err != nil {
		logger.Error("Error to delete seller", journey, err)
		return err
	}

	return nil
}
