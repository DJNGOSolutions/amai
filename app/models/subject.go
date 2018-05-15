package models

import "github.com/jinzhu/gorm"

type Subject struct {			//Materia
	gorm.Model
	idSubject     string `gorm:"primary_key"`		//idMateria
	SubjectName string `gorm:"type:varchar(75);unique"`	//nombreMateria
}