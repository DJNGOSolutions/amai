package models

import "github.com/jinzhu/gorm"

type Cita struct{

	gorm.Model

	IdCita uint `gorm: "primary key"`

	IdLugarCita uint	//fk vinculada con el catálogo de lugares
	IdMateria uint	//fk vinculada con la tabla de Materias
	IdTema uint		//fk vinculada con la tabla Temas
	idTipoCita uint		//fk vinculada con el catálogo de tipo de cita
	idTipoRemuneracion uint //fk vinculada con el catálogo de tipo de remuneracion

	FechaCita string `gorm: "type: date"`
	HoraInicioCita string `gorm: "type: time"`
	HoraFinalizacionCita string `gorm: "type: time"`
	PrecioCita string `gorm: "type: money"`

}
