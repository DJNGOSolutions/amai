package models

import "github.com/jinzhu/gorm"

type Tutoria struct {
	gorm.Model
	IdTutoria uint `gorm:"primary_key"`
	IdUsuario uint
	IdMateria uint
}