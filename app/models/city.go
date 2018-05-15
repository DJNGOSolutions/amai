package models

import "github.com/jinzhu/gorm"

type City struct {				//Municipio
	gorm.Model
	IdCity     uint `gorm:"primary_key"`			//idMunicipio
	IdState  uint								//idDepartamento
	CityName string `gorm:"type:varchar(100)"`	//nombreMunicipio
}
