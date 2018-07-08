package models

import "github.com/jinzhu/gorm"

type Place_Session struct {
	//Lugar_Session
	//Esta estructura es un catálogo
	gorm.Model
	Place string `gorm:"type:varchar(100)"` //idLugar
}
