package models

import "github.com/jinzhu/gorm"

type State struct {
	//Departamento
	gorm.Model
	IdState     uint   `gorm:"primary_key"` 		//idDepartamento
	StateName 	string `gorm:"type:varchar(100)"`  	//NombreDepartamento
}
