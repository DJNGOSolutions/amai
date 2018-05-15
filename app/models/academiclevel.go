package models

import "github.com/jinzhu/gorm"

type AcademicLevel_User struct {					//nivelAcademico_Usuario
	//Esta estructura es un catálogo
	gorm.Model
	IdAcademicLevel_User uint   `gorm:"primary_key"`	//idNivelEducativo_Usuario
	AcademicaLevel           string `gorm:"type:varchar(100);unique"` //nivelEducativo
}
