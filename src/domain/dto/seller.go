package dto

import (
	"github.com/pedro-phd/iservice-backend/src/domain/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SellerRequestDTO struct {
	Name     string           `json:"name" binding:"required,min=3,max=255"`
	Email    string           `json:"email" binding:"required,email"`
	Password string           `json:"password" binding:"required,min=8,max=64,containsany=!@#%$_."`
	Products []models.Product `json:"products"`
}

type SellerResponseDTO struct {
	ID       primitive.ObjectID   `json:"id"`
	Name     string               `json:"name"`
	Email    string               `json:"email"`
	password string               `json:"password"`
	Products []ProductResponseDTO `json:"products"`
}

func NewSellerMongoDTO(name, email, password string, products []models.Product) SellerMongoDTO {

	prodDTO := make([]ProductMongoDTO, 0)

	for _, prod := range products {
		temp := NewProductToMongo(prod.GetName(), prod.GetPrice(), prod.GetDescription(), prod.GetThumb())
		prodDTO = append(prodDTO, temp)
	}

	return SellerMongoDTO{
		Name:     name,
		Email:    email,
		Password: password,
		Products: prodDTO,
	}
}

func NewMongoToSeller(name, email, password string, products []models.Product) models.Seller {
	return models.NewSeller(name, email, password, products)
}

type SellerMongoDTO struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Products []ProductMongoDTO  `bson:"products" json:"products"`
}
