package models

import "github.com/jinzhu/gorm"

type Institucion struct {
	gorm.Model
	IdInstitucion     uint `gorm:"primary_key"`
	IdMateria         uint
	idUsuario         uint
	idDepartamento    uint
	idMunicipio       uint
	NombreInstitucion string `gorm:"type:varchar(150)"`
}
