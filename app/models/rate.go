package models

import "github.com/jinzhu/gorm"

type Rate_User struct {
	//Esta estructura es un catálogo
	gorm.Model
	Rate string `gorm:"type:varchar(20);unique"` //Rol
}
