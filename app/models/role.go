package models

import "github.com/jinzhu/gorm"

type Role_User struct {
	//Esta estructura es un catálogo
	gorm.Model
	IdRole_User uint   `gorm:"primary_key"`		//idRole_User
	Role           string `gorm:"type:varchar(100);unique"` //Rol
}
