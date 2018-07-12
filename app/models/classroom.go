package models

import "github.com/jinzhu/gorm"

type Classroom struct {
	gorm.Model
	IdCategory      uint   `gorm:"type:integer"`
	IdSubject       uint   `gorm:"type:integer"`
	IdUserTutor     uint   `gorm:"type:integer"`
	IdUserCustomer  uint   `gorm:"type:integer"`
	ClassroomPhoto  string `gorm:"type:text"`
	ClassroomTheme  string `gorm:"type:varchar(30)"`
	ClassroomPlace  string `gorm:"type:text"`
	ClassroomTime   string `gorm:"type:time"`
	ClassroomDate   string `gorm:"type:date"`
	ClassroomPrince string `gorm:"type:money"`
}
