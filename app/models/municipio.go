package models

import "github.com/jinzhu/gorm"

type Municipio struct{

	gorm.Model

	IdMunicipio uint `gorm: "primary key"`
	IdDepartamento uint

	NombreMunicipio string `gorm: "type:varchar(100)"`

}
