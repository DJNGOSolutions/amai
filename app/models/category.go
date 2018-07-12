package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CategoryPhoto string `gorm:"type:text"`
<<<<<<< HEAD
	CategoryName  string `gorm:"type:varchar(40),unique"`
=======
	CategoryName  string `gorm:"type:varchar(100);unique"`
>>>>>>> 6dfec73d16e44dbc24abc6260d18bb75e4a5f5cc
}
