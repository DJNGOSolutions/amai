package models

import "github.com/jinzhu/gorm"

type RatexUser struct {
	//Esta estructura es un catálogo
	gorm.Model
	IdUser uint //idUsuario
	Rate   float32 //Rol
}
