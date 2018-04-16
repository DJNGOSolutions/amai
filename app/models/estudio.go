package models

import "github.com/jinzhu/gorm"

type Estudio struct {
	gorm.Model
	IdEstudio uint `gorm:"primary_key"`
	IdUsuario uint
	IdMateria uint
}
