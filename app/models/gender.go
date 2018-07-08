package models

import "github.com/jinzhu/gorm"

type Gender_User struct { //Genero_Usuario
	//Esta estructura es un catálogo
	gorm.Model
	Gender string `gorm:"type:varchar(100);unique"` //Genero
}
