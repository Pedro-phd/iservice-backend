package models

type Product interface {
	GetName() string
	GetPrice() float64
	GetDescription() string
	GetThumb() string
}

type product struct {
	name        string
	price       float64
	description string
	thumb       string
}

func (p *product) GetName() string        { return p.name }
func (p *product) GetPrice() float64      { return p.price }
func (p *product) GetDescription() string { return p.description }
func (p *product) GetThumb() string       { return p.thumb }

func NewProduct(name string, price float64, description string, thumb string) Product {
	return &product{name, price, description, thumb}
}
