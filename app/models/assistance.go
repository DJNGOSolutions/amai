package models

import "github.com/jinzhu/gorm"

<<<<<<< HEAD
type Assistance struct { //Asistencia
	gorm.Model
	IdUser          uint   //idUsuario
	IdSession       uint   //idCita
	TotalFeeSession string `gorm:"type:money"` //¿Cómo hacer un atributo derivado? - PrecioTotalAsistencia
=======
type Assistance struct {
	//Asistencia
	gorm.Model
	IdAssistance      	uint 	`gorm:"primary_key"` 	//idAsistencia
	IdUser             	uint 							//idUsuario
	IdSession         	uint 							//idCita
	TotalFeeSession 	string 	`gorm:"type:money"` 	//¿Cómo hacer un atributo derivado? - PrecioTotalAsistencia
>>>>>>> 6734364bcb548ac850daea73afb8769d9a84f9c5
}
