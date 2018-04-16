package models

import "github.com/jinzhu/gorm"

type Usuario struct{

	gorm.Model

	IdUsuario uint `gorm: "primary key"`

	IdGeneroUsuario uint //fk vinculada con el catálogo de género
	IdNivelEducativoUsuario uint //fk vinculada con el catálogo de nivel educativo
	IdRolUsuario uint //fk vinculada con el catálogo de rol

	NombreUsuario string `gorm: "type: varchar(150);not null"`
	EdadUsuario uint8
	correoUsuario string `gorm: "type: varchar(100);not null;unique"`
	telefonoUsuario string `gorm "type: varchar(30);unique"` //multivaluado - cambiar

}