package models

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	ProductID uint `gorm:"primary_key"`
	Code      string
	Price     uint
}
