package models

import "github.com/jinzhu/gorm"

type AcademicLevel_User struct { //nivelAcademico_Usuario
	//Esta estructura es un catálogo
	gorm.Model
	AcademicLevel string `gorm:"type:varchar(100);unique"` //nivelEducativo
}
