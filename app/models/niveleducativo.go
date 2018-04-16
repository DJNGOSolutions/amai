package models

import "github.com/jinzhu/gorm"

type NivelEducativo_Usuario struct{
	//Esta estructura es un catálogo
	gorm.Model

	IdNivelEducativo_Usuario uint `gorm: "primary key"`

	NivelEducativo string `gorm: "type: varchar(100);unique"`

}