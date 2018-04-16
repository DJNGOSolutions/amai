package models

import "github.com/jinzhu/gorm"

type TipoRemuneracion_Cita struct {
	//Esta estructura es un catálogo
	gorm.Model
	IdTipoRemuneracion_Cita uint   `gorm:"primary_key"`
	TipoRemuneracion        string `gorm:"type:varchar(100);unique"`
}
