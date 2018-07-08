package models

import "github.com/jinzhu/gorm"

type SubjectxTopic struct {
	//MateriaxTema
	gorm.Model
	IdSubject uint
	IdTopic   uint
}
