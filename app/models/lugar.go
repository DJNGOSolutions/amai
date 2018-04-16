package models

import "github.com/jinzhu/gorm"

type Lugar_Cita struct{
	//Esta estructura es un catálogo
	gorm.Model

	IdLugar_Cita uint `gorm: "primary key"`

	Lugar string `gorm: "type: varchar(100)"`

}
