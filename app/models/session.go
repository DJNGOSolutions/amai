package models

import "github.com/jinzhu/gorm"

type Session struct { //Cita
	gorm.Model
	IdSessionPlace   uint   //fk vinculada con el catálogo de lugares - idLugarSession
	IdSubject        uint   //fk vinculada con la tabla de Materias - idMateria
	IdTopic          uint   //fk vinculada con la tabla Temas - idTopic
	IdTypeSession    uint   //fk vinculada con el catálogo de tipo de cita - idTipoCita
	IdTypePayment    uint   //fk vinculada con el catálogo de tipo de remuneracion - idTipoRemuneración
	SessionDate      string `gorm:"type:date"`  //FechaCita
	SessionTimeStart string `gorm:"type:time"`  //horaInicioCita
	SessionTimeEnd   string `gorm:"type:time"`  //horaFinalizacionCita
	SessionFee       string `gorm:"type:money"` //precioCita
}
