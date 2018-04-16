package models

import "github.com/jinzhu/gorm"

type Rol_Usuario struct {
	//Esta estructura es un catálogo
	gorm.Model
	IdRol_Usuario uint   `gorm:"primary_key"`
	Rol           string `gorm:"type:varchar(100);unique"`
}
