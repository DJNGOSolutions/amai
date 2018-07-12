package models

import "github.com/jinzhu/gorm"

type Subject struct {
	//Materia
	gorm.Model
	IdCategory  uint   `gorm:"tye:varchar(75)"`
	SubjectName string `gorm:"type:varchar(75);unique"`
}
