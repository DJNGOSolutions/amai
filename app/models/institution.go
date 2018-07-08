package models

import "github.com/jinzhu/gorm"

type Institution struct { //Institution
	gorm.Model
	IdSubject       uint   //idMateria
	IdUser          uint   //idUsuario
	IdState         uint   //idDepartamento
	IdCity          uint   //idMunicipio
	InstitutionName string `gorm:"type:varchar(150)"` //nombreInstitucion
}
