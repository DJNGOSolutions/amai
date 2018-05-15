package models

import "github.com/jinzhu/gorm"

type Tutorial struct {					//Tutoria
	gorm.Model
	IdTutorial uint `gorm:"primary_key"`	//idTutoria
	IdUser uint								//idUsuario
	IdSubject uint							//idMateria
}
