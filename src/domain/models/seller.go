package models

import (
	"crypto/md5"
	"encoding/hex"
)

type Seller interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetProducts() []Product

	EncryptPassword()
}

type seller struct {
	name     string
	email    string
	password string
	products []Product
}

func (s *seller) GetName() string        { return s.name }
func (s *seller) GetEmail() string       { return s.email }
func (s *seller) GetPassword() string    { return s.password }
func (s *seller) GetProducts() []Product { return s.products }

func (s *seller) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(s.password))
	s.password = hex.EncodeToString(hash.Sum(nil))
}

func NewSeller(name, email, password string, products []Product) Seller {
	return &seller{name, email, password, products}
}
