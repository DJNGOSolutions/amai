package models

import "github.com/jinzhu/gorm"

type PhonexUser struct{					//TelefonoxUsuario
	//Esta es una tabla multivaluada
	gorm.Model
	IdUser uint						//idUsuario
	IdPhonexUser uint `gorm:"primary_key"` //idTelefonoxUsuario
	Phone string `gorm:"type:varchar(30);unique"`	//Telefono
}
