package models

import "github.com/jinzhu/gorm"

type User struct {				//Usuario
	gorm.Model
	IdUser              uint   `gorm:"primary_key"`		//idUsuario
	IdGenderUser         uint   //fk vinculada con el catálogo de género - idGeneroUsuario
	IdAcademicLevelUser uint   //fk vinculada con el catálogo de nivel educativo - idNivelEducativoUsuario
	IdRoleUser            uint   //fk vinculada con el catálogo de rol - idRolUsuario
	UserName           string `gorm:"type:varchar(150);not null"` //nombreUsuario
	UserAge             uint8 //edadUsuario
	UserEmail           string `gorm:"type:varchar(100);not null;unique"` //correoUsuario
	UserPhone         string `gorm:"type:varchar(30);unique"` //multivaluado - cambiar - telefonoUsuario
	UserDescription		string `gorm:"type:text"`
}