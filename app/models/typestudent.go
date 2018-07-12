package models

import "github.com/jinzhu/gorm"

type TypeStudent struct {
	gorm.Model
	TypeName string `gorm:"type:varchar(40);unique"`
}
