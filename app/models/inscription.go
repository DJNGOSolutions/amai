package models

import "github.com/jinzhu/gorm"

type Inscription struct {
	//Asistencia
	gorm.Model
	IdUser      int `gorm:"type:integer"`
	IdClassroom int `gorm:"type:integer"`
}
