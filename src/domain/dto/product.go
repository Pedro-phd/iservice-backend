package dto

import "time"

type ProductRequestDTO struct {
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required,min=0"`
	Description string  `json:"description" binding:"required"`
	Thumb       string  `json:"thumb" binding:"required"`
}

type ProductResponseDTO struct {
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	Thumb       string    `json:"thumb"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewProductToMongo(name string, price float64, description, thumb string) ProductMongoDTO {
	return ProductMongoDTO{Name: name, Price: price, Description: description, Thumb: thumb}
}

type ProductMongoDTO struct {
	Name        string    `bson:"name,omitempty" json:"name,omitempty"`
	Price       float64   `bson:"price,omitempty" json:"price,omitempty"`
	Description string    `bson:"description,omitempty" json:"description,omitempty"`
	Thumb       string    `bson:"thumb,omitempty" json:"thumb,omitempty"`
	CreatedAt   time.Time `bson:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at"`
}
