package models

import "github.com/jinzhu/gorm"

type Session struct { //Cita
	gorm.Model
	IdCategory     uint   `gorm:"type:integer"`
	IdSubject      uint   `gorm:"type:integer"`
	IdUserTutor    uint   `gorm:"type:integer"`
	IdUserCustomer uint   `gorm:"type:integer"`
	TutoryPhoto    string `gorm:"type:text"`
	TutoryTheme    string `gorm:"type:varchar(30)"`
	TutoryPlace    string `gorm:"type:text"`
	TutoryTime     string `gorm:"type:time"`
	TutoryDate     string `gorm:"type:date"`
	TutoryPrice    string `gorm:"type:money"`
}
