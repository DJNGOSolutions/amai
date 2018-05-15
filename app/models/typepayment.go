package models

import "github.com/jinzhu/gorm"

type TypePayment_Session struct {
	//Esta estructura es un catálogo
	gorm.Model
	IdTypePayment_Session uint   `gorm:"primary_key"`
	TypePayment        string `gorm:"type:varchar(100);unique"`
}
