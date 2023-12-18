package dto

import (
	"github.com/pedro-phd/iservice-backend/src/domain/models"
)

type SellerRequest struct {
	Name       string `json:"name,omitempty" db:"Name"`
	Email      string `json:"email,omitempty" db:"Email"`
	Document   string `json:"document,omitempty" db:"Document"`
	Phone      string `json:"phone,omitempty" db:"Phone"`
	Password   string `json:"password,omitempty" db:"Password"`
	Image      string `json:"image,omitempty" db:"Image"`
	State      string `json:"state,omitempty" db:"State"`
	City       string `json:"city,omitempty" db:"City"`
	Street     string `json:"street,omitempty" db:"Street"`
	Number     string `json:"number,omitempty" db:"Number"`
	Active     bool   `json:"active,omitempty" db:"Active"`
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
