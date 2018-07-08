package models

import "github.com/jinzhu/gorm"

type Study struct {
	//Estudio
	gorm.Model
	IdUser    uint //idUsuario
	IdSubject uint //idMateria
}
