package models

import "github.com/jinzhu/gorm"

type Institucion struct{

	gorm.Model

	IdInstitucion uint `gorm: "primary key"`
	IdMateriaInstitucion uint
	idUsuario uint

	NombreInstitucion string `gorm: "type: varchar(150)"`
	//PaisInstitucion string `gorm: "type: varchar(100)"`
	DepartamentoInstitucion string `gorm: "type: varhcar(100) unique"`  //catálogo - cambiar
	MunicipioInstitucion string `gorm: "type: varchar(100); unique"`	//catálogo - cambiar

}
