package models

import "github.com/jinzhu/gorm"

type UserT struct {
	//Usuario
	gorm.Model
	UserPhoto           string 	`gorm:"type:text"`                  //fotoUsuario
	UserName            string 	`gorm:"type:varchar(150);not null"` //nombreUsuario
	UserInstitution		string	`gorm:"type:varchar(50)"`
	UserPassword        string 	`gorm:"type:text"`
	IdTipo				uint
}
