package models

import "github.com/jinzhu/gorm"

type Session struct { //Cita
	gorm.Model
	IdCategory		uint
	IdSubject		uint
	IdUserTutor		uint
	IdUserCustomer	uint
	TutoryPhoto	string	`gorm:"type:text"`
	TutoryTheme	string	`gorm:"type:varchar(30)"`
	TutoryPlace	string	`gorm:"type:text"`
	TutoryTime	string	`gorm:"type:time"`
	TutoryDate	string	`gorm:"type:date"`
	TutoryPrince	string	`gorm:"type:money"`
}
