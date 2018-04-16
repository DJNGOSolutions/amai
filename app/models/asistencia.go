package models

import "github.com/jinzhu/gorm"

type Asistencia struct {
	gorm.Model
	IdAsistencia          uint `gorm:"primary_key"`
	IdUsuario             uint //¿Qué pasa si a esta tabla se pasan hasta más de una fk idUsuario? - Varela
	IdCita                uint
	PrecioTotalAsistencia string `gorm:"type:money"` //¿Cómo hacer un atributo derivado?
}
