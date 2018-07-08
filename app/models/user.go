package models

import "github.com/jinzhu/gorm"

type User struct { //Usuario
	gorm.Model
	IdGenderUser        uint   //fk vinculada con el catálogo de género - idGeneroUsuario
	IdAcademicLevelUser uint   //fk vinculada con el catálogo de nivel educativo - idNivelEducativoUsuario
	IdRoleUser          uint   //fk vinculada con el catálogo de rol - idRolUsuario
	IdRateUser          uint   //fk vinculada con el catálogo de rate - idCalificacionUsuario
	UserPhoto           string `gorm:"type:text"`                  //fotoUsuario
	UserName            string `gorm:"type:varchar(150);not null"` //nombreUsuario
	UserAge             uint8  //edadUsuario,f
	UserEmail           string `gorm:"type:varchar(100);not null;unique"` //correoUsuario
	UserDescription     string `gorm:"type:text"`                         //descripcionUsuario
}
