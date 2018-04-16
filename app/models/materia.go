package models

import "github.com/jinzhu/gorm"

type Materia struct{

	gorm.Model

	idMateria string `gorm: "primary key"`

	NombreMateria string `gorm: "type: varchar(75); unique"`

}
