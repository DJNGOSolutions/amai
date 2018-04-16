package models

import "github.com/jinzhu/gorm"

type Tema struct {
	gorm.Model
	idTema     uint   `gorm:"primary_key"`
	NombreTema string `gorm:"type:varchar(75);unique"`
}
