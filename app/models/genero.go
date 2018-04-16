package models

import "github.com/jinzhu/gorm"

type Genero_Usuario struct{
	//Esta estructura es un catálogo
	gorm.Model

	IdGenero_Usuario uint `gorm: "primary key"`

	Genero string `gorm: "type: varchar(100), unique"`

}
