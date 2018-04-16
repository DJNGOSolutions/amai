package models

import"github.com/jinzhu/gorm"

type Clase struct{

	gorm.Model

	IdClase uint `gorm: "primary key"`
	IdInstitucion uint
	IdMateria uint

}
