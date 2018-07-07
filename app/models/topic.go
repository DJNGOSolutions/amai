package models

import "github.com/jinzhu/gorm"

type Topic struct {
	//Tema
	gorm.Model
	IdTopic     uint   `gorm:"primary_key"`				//idTopic
	TopicName 	string `gorm:"type:varchar(75);unique"`	//NombreTema
}
