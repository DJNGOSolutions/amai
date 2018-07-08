package models

import "github.com/jinzhu/gorm"

type TypePayment_Session struct {
	//Esta estructura es un catálogo
	gorm.Model
	TypePayment string `gorm:"type:varchar(100);unique"`
}
