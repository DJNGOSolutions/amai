package models

import "github.com/jinzhu/gorm"

type MateriaxTema struct {
	gorm.Model
	IdMateriaxTema uint `gorm:"primary_key"`
	IdMateria      uint
	IdTema         uint
}
