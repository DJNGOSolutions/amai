package models

import "github.com/jinzhu/gorm"

type UserxClassroom struct {
	gorm.Model
	idUser			uint
	idClassroom		uint
}