package models

import "github.com/jinzhu/gorm"

type Classroom struct{
	gorm.Model
	IdSubject	uint
	IdTopic		uint
}
