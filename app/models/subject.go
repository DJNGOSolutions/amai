package models

import "github.com/jinzhu/gorm"

type Subject struct { //Materia
	gorm.Model
	SubjectName string `gorm:"type:varchar(75);unique"` //nombreMateria
}
