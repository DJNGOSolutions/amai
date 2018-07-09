package models

import "github.com/jinzhu/gorm"

type Topic struct {
	//Tema
	gorm.Model
	IdSubject	uint
	TopicName	string `gorm:"type:varchar(75);unique"` //NombreTema
}
