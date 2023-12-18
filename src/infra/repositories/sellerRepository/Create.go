package sellerRepository

import (
	"fmt"
	"strconv"

	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
)

func (sr *SellerRepository) Create(sellerRequest *dto.SellerRequest) (string, error) {
	journey := "Create - Seller"

	logger.Info("Init to create seller", journey)

	seller := dto.MapRequestToSeller(*sellerRequest)

	seller.EncryptPassword()

	sellerResult, err := sr.db.NamedQuery(`
		INSERT INTO sellers 
		(name, email, document, phone, password, image, state, city, street, number, active)
		VALUES
		(:Name, :Email, :Document, :Phone, :Password, :Image, :State, :City,:Street,:Number, true)
		RETURNING id
	`, seller)

	if err != nil {
		logger.Error("Error to create seller", journey, err)
		return "", err
	}

	var id int64
	if sellerResult.Next() {
		sellerResult.Scan(&id)
	}

	for _, categoryId := range sellerRequest.Categories {
		_, err := sr.db.Exec(`
			INSERT INTO sellers_categories 
			(seller_id, category_id)
			VALUES
			($1, $2)
			`, id, categoryId)

		if err != nil {
			errMsg := fmt.Sprintf("Error to insert seller in correct category, category:%d", categoryId)
			logger.Error(errMsg, journey, err)
			return "", err
		}

	}

	idStrg := strconv.FormatInt(id, 10)
	return idStrg, nil
}
