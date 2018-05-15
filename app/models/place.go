package models

import "github.com/jinzhu/gorm"

type Place_Session struct {				//Lugar_Session
	//Esta estructura es un catálogo
	gorm.Model
	IdPlace_Session uint   `gorm:"primary_key"`		//idLugar_Session
	Place        string `gorm:"type:varchar(100)"`	//idLugar
}
