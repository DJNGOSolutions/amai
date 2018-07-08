package models

import "github.com/jinzhu/gorm"

type Class struct { //Clase
	gorm.Model
	IdInstitution uint //idInstitucion
	IdSubject     uint //idMateria
}
