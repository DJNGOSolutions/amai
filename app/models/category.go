package models

import "github.com/jinzhu/gorm"

type category struct{
	gorm.Model
	CategoryPhoto	string `gorm:type:text`
	CategoryName	string	`gorm:"type:varchar(100),unique"`
}
