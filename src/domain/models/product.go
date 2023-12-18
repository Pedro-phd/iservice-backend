package models

type Product struct {
	BasicModel
	SellerId    int
	Name        string
	Price       float64
	Description string
	Image       string
	Active      bool
}
