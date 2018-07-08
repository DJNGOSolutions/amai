package models

import "github.com/jinzhu/gorm"

type Tutorial struct {
	//Tutoria
	gorm.Model
	IdUser    uint //idUsuario
	IdSubject uint //idMateria
}
