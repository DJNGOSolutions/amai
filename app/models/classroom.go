package models

import "github.com/jinzhu/gorm"

type Classroom struct{
	gorm.Model
	IdTopic		uint
	IdSubject	uint
}
