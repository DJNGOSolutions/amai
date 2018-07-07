package models

import "github.com/jinzhu/gorm"

type RatexUser struct {
	//Esta estructura es un catálogo
	gorm.Model
	IdRatexUser 	uint   	`gorm:"primary_key"`				//idRole_User
	IdUser 			uint										//idUsuario
	Rate           	string 	`gorm:"type:varchar(20);unique"` 	//Rol
}