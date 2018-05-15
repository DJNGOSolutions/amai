package models

import "github.com/jinzhu/gorm"

type Study struct {				//Estudio
	gorm.Model
	IdStudy uint `gorm:"primary_key"`	//idEstudio
	IdUser uint							//idUsuario
	IdSubject uint						//idMateria
}
