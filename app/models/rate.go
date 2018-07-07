package models

import "github.com/jinzhu/gorm"

type Rate_User struct {
	//Esta estructura es un catálogo
	gorm.Model
	IdRate_User 	uint   `gorm:"primary_key"`				//idRole_User
	Rate           	string `gorm:"type:varchar(20);unique"` //Rol
}