package models

type Order struct {
	BasicModel
	SellerId int
	ClientId int
	Status   string
}
