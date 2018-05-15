package models

import "github.com/jinzhu/gorm"

type SubjectxTopic struct {				//MateriaxTema

	gorm.Model
	IdSubjectXTopc uint `gorm:"primary_key"`
	IdSubject      uint
	IdTopic         uint
}
