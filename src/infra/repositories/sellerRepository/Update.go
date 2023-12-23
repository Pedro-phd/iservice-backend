package sellerRepository

import (
	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
)

func (sr *SellerRepository) Update(id int64, seller *dto.SellerRequest) error {

	journey := "Update - Seller"
	logger.Info("Init to update seller", journey)

	seller.Id = id

	_, err := sr.db.NamedExec(`
	UPDATE sellers SET
		name = CASE WHEN :Name <> '' AND :Name <> name THEN :Name ELSE name END,
		email = CASE WHEN :Email <> '' AND :Email <> email THEN :Email ELSE email END,
		document = CASE WHEN :Document <> '' AND :Document <> document THEN :Document ELSE document END,
		phone = CASE WHEN :Phone <> '' AND :Phone <> phone THEN :Phone ELSE phone END,
		password = CASE WHEN :Password <> '' THEN :Password ELSE password END,
		image = CASE WHEN :Image <> '' THEN :Image ELSE image END,
		state = CASE WHEN :State <> '' THEN :State ELSE state END,
		city = CASE WHEN :City <> '' THEN :City ELSE city END,
		street = CASE WHEN :Street <> '' THEN :Street ELSE street END,
		number = CASE WHEN :Number <> '' THEN :Number ELSE number END,
		active = COALESCE(:Active, active)
	WHERE id = :Id
	`, seller)

	if err != nil {
		logger.Error("Error to update seller", journey, err)
	}

	return nil
}
