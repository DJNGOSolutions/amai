package models

import "github.com/jinzhu/gorm"

type Clase struct {
	gorm.Model
	IdClase       uint `gorm:"primary_key"`
	IdInstitucion uint
	IdMateria     uint
}
