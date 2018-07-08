package models

import "github.com/jinzhu/gorm"

type Assistance struct {
	//Asistencia
	gorm.Model
	IdUser          uint   //idUsuario
	IdSession       uint   //idCita
	TotalFeeSession string `gorm:"type:money"` //¿Cómo hacer un atributo derivado? - PrecioTotalAsistencia
}
