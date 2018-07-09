package models

import "github.com/jinzhu/gorm"

type Subscription struct {
	//Tabla multivaluada
	gorm.Model
	IdClassroom		uint
	IdUser			uint
}
