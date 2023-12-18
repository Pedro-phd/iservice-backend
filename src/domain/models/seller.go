package models

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type Seller struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty" db:"Name"`
	Email     string    `json:"email,omitempty" db:"Email"`
	Document  string    `json:"document,omitempty" db:"Document"`
	Phone     string    `json:"phone,omitempty" db:"Phone"`
	Password  string    `json:"password,omitempty" db:"Password"`
	Image     string    `json:"image,omitempty" db:"Image"`
	State     string    `json:"state,omitempty" db:"State"`
	City      string    `json:"city,omitempty" db:"City"`
	Street    string    `json:"street,omitempty" db:"Street"`
	Number    string    `json:"number,omitempty" db:"Number"`
	Active    bool      `json:"active,omitempty" db:"Active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ISeller interface {
	EncryptPassword()
}

func (s *Seller) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(s.Password))
	s.Password = hex.EncodeToString(hash.Sum(nil))
}
