package models

import "github.com/jinzhu/gorm"

type Tipo_Cita struct {
	//Esta estructura es un catálogo
	gorm.Model
	IdTipo_Cita uint   `gorm:"primary_key"`
	TipoCita    string `gorm:"type:varchar(100)"`
}
