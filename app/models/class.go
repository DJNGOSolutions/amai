package models

import "github.com/jinzhu/gorm"

type Class struct {				//Clase
	gorm.Model
	IdClass       uint `gorm:"primary_key"` //idClase
	IdInstitution uint //idInstitucion
	IdSubject     uint //idMateria
}
