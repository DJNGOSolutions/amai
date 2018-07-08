package models

import "github.com/jinzhu/gorm"

type State struct {
	//Departamento
	gorm.Model
	StateName string `gorm:"type:varchar(100)"` //NombreDepartamento
}
