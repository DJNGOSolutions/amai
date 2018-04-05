package models

type Product struct {
	ProductID uint `gorm:"primary_key"`
	Code      string
	Price     uint
}
