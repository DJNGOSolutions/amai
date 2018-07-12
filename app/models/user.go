package models

import "github.com/jinzhu/gorm"

type UserT struct {
	//Usuario
	gorm.Model
	UserPhoto       string `gorm:"type:text"`                  //fotoUsuario
	UserName        string `gorm:"type:varchar(150);not null"` //nombreUsuario
	UserInstitution string `gorm:"type:varchar(50)"`
	UserPassword    []byte `gorm:"type:text"`
	IdGoogle        uint   `gorm:"type:integer;unique"`
	IdTipo          uint   `gorm:"type:integer"`
}
