package models

type SellersCategories struct {
	Id         uint `json:"id,omitempty" db:"id"`
	SellerID   uint `json:"seller_id,omitempty" db:"seller_id"`
	CategoryID uint `json:"category_id,omitempty" db:"category_id"`
}
