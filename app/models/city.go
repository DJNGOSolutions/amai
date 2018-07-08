package models

import "github.com/jinzhu/gorm"

type City struct { //Municipio
	gorm.Model
	IdState  uint   //idDepartamento
	CityName string `gorm:"type:varchar(100)"` //nombreMunicipio
}
