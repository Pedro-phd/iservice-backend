package sellerRepository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
)

type ISellerRepository interface {
	Create(sellerRequest *dto.SellerRequest) (string, error)
	// FindById(sellerId int) (*models.Seller, error)
	// FindAll() ([]*models.Seller, error)
	// Update(seller models.Seller) error
	Delete(sellerId int) error
}

type SellerRepository struct {
	db *sqlx.DB
}

func NewSellerRepository(db *sqlx.DB) ISellerRepository {
	return &SellerRepository{db: db}
}
