package sellerRepository

import (
	"fmt"

	"github.com/pedro-phd/iservice-backend/src/configuration/logger"
	"github.com/pedro-phd/iservice-backend/src/domain/models"
)

func (sr *SellerRepository) Delete(sellerId int) error {

	journey := "Delete - Seller"
	logger.Info("Init to delete seller", journey)

	var categoriesResult []models.SellersCategories
	err := sr.db.Select(&categoriesResult, "SELECT * FROM sellers_categories WHERE seller_id = $1", sellerId)

	if err != nil {
		logger.Error("Error to query all categories for this seller", journey, err)
	}

	// if categories.Next() {
	// 	categories.Scan(&categoriesResult)
	// }
	for _, category := range categoriesResult {
		fmt.Println(category.Id)
	}
	// fmt.Println(categoriesResult)
	// for _, cat := range categoriesResult {
	// 	fmt.Println(cat)
	// }
	// fmt.Println(categories.Next(), categoriesResult, sellerId)

	// _, err := sr.db.Exec("DELETE FROM sellers WHERE id = $1", sellerId)
	// if err != nil {
	// 	logger.Error("Error to delete seller", journey, err)
	// 	return err
	// }

	return nil
}
