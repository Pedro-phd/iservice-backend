package sellerrepository

import (
	"github.com/pedro-phd/iservice-backend/src/domain/dto"
	rest_err "github.com/pedro-phd/iservice-backend/src/domain/error"
	"github.com/pedro-phd/iservice-backend/src/domain/models"
)

type ISellerRepository interface {
	GetAll() ([]dto.SellerResponseDTO, *rest_err.RestErr)
	// getByID(string) models.Seller
	Create(models.Seller) (string, *rest_err.RestErr)
	Remove(string) bool
}

type SellerRepository struct{}

func NewSellerRepository() ISellerRepository {
	return &SellerRepository{}
}
