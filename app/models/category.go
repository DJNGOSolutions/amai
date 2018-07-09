package models

import "github.com/jinzhu/gorm"

type Category struct{
	gorm.Model
	CategoryPhoto	string `gorm:type:text`
	CategoryName	string	`gorm:"type:varchar(100),unique"`
}
