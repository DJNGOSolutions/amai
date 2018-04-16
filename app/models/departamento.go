package models

import "github.com/jinzhu/gorm"

type Departamento struct{

	gorm.Model

	IdDepartamento uint `gorm: "primary key"`

	NombreDepartamento string `gorm: "type:varchar(100)"`
}
