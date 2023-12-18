package models

type Categorie struct {
	BasicModel
	Name string `gorm:"unique;not null"`
}
