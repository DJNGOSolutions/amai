package models

import "github.com/jinzhu/gorm"

type Subject struct {
	//Materia
	gorm.Model
	IdCategory	uint
	SubjectName string `gorm:"type:varchar(75);unique"` //nombreMateria
}
