package models

import"github.com/jinzhu/gorm"

type MateriaxTema struct{

	gorm.Model

	IdMateriaxTema uint `gorm: "primary key"`
	IdMateria uint
	IdTema uint

}
