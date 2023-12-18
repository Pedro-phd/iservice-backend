package models

type OrdersItems struct {
	BasicModel
	OrderID   int
	ProductID int
	Quantity  int
	Price     float64
}
