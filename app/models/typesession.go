package models

import "github.com/jinzhu/gorm"

type Type_Session struct {
	//Tipo_Cita
	//Esta estructura es un catálogo
	gorm.Model
	TypeSession string `gorm:"type:varchar(100)"`
}
