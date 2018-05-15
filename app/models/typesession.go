package models

import "github.com/jinzhu/gorm"

type Type_Session struct {				//Tipo_Cita
	//Esta estructura es un catálogo
	gorm.Model
	IdType_Session uint   `gorm:"primary_key"`
	TypeSession    string `gorm:"type:varchar(100)"`
}
