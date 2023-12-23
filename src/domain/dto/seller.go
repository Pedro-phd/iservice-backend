package dto

import (
	"github.com/pedro-phd/iservice-backend/src/domain/models"
)

type SellerResponse struct {
	Id         int64                      `json:"id,omitempty" db:"id"`
	Name       string                     `json:"name,omitempty" db:"name"`
	Email      string                     `json:"email,omitempty" db:"email"`
	Document   string                     `json:"document,omitempty" db:"document"`
	Phone      string                     `json:"phone,omitempty" db:"phone"`
	Password   string                     `json:"-" db:"password"`
	Image      string                     `json:"image,omitempty" db:"image"`
	State      string                     `json:"state,omitempty" db:"state"`
	City       string                     `json:"city,omitempty" db:"city"`
	Street     string                     `json:"street,omitempty" db:"street"`
	Number     string                     `json:"number,omitempty" db:"number"`
	Active     bool                       `json:"active,omitempty" db:"active"`
	Categories []models.SellersCategories `json:"categories"`
}

type SellerRequest struct {
	Id         int64  `json:"id,omitempty" db:"Id"`
	Name       string `json:"name,omitempty" db:"Name" binding:"required,min=3"`
	Email      string `json:"email,omitempty" db:"Email" binding:"required,email"`
	Document   string `json:"document,omitempty" db:"Document" binding:"required,min=11,excludesrune=.,excludesrune=-"`
	Phone      string `json:"phone,omitempty" db:"Phone" binding:"required,min=9,excludesrune=.,excludesrune=-,excludesrune=(,excludesrune=)"`
	Password   string `json:"password,omitempty" db:"Password" binding:"required,min=8,containsany=!@#%$_.`
	Image      string `json:"image,omitempty" db:"Image" binding:"required"`
	State      string `json:"state,omitempty" db:"State" binding:"required"`
	City       string `json:"city,omitempty" db:"City" binding:"required"`
	Street     string `json:"street,omitempty" db:"Street" binding:"required"`
	Number     string `json:"number,omitempty" db:"Number" binding:"required"`
	Active     bool   `json:"active,omitempty" db:"Active" binding:"required,boolean"`
	Categories []int64
}

func MapRequestToSeller(sr SellerRequest) models.Seller {
	return models.Seller{
		Name:     sr.Name,
		Email:    sr.Email,
		Document: sr.Document,
		Phone:    sr.Phone,
		Password: sr.Password,
		Image:    sr.Image,
		State:    sr.State,
		City:     sr.City,
		Street:   sr.Street,
		Number:   sr.Number,
		Active:   sr.Active,
	}
}
