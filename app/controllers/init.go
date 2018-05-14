package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(InitDB)
}

func InitDB() {
	revel.INFO.Println("Doing DB Migrations...")
	/*
		gormdb.DB.AutoMigrate(&models.Asistencia{})
		gormdb.DB.AutoMigrate(&models.Cita{})
		gormdb.DB.AutoMigrate(&models.Clase{})
		gormdb.DB.AutoMigrate(&models.Departamento{})
		gormdb.DB.AutoMigrate(&models.Estudio{})
		gormdb.DB.AutoMigrate(&models.Genero_Usuario{})
		gormdb.DB.AutoMigrate(&models.Institucion{})
		gormdb.DB.AutoMigrate(&models.Lugar_Cita{})
		gormdb.DB.AutoMigrate(&models.Materia{})
		gormdb.DB.AutoMigrate(&models.MateriaxTema{})
		gormdb.DB.AutoMigrate(&models.Municipio{})
		gormdb.DB.AutoMigrate(&models.NivelEducativo_Usuario{})
		gormdb.DB.AutoMigrate(&models.Rol_Usuario{})
		gormdb.DB.AutoMigrate(&models.Tema{})
		gormdb.DB.AutoMigrate(&models.Tipo_Cita{})
		gormdb.DB.AutoMigrate(&models.TipoRemuneracion_Cita{})
		gormdb.DB.AutoMigrate(&models.Tutoria{})
		gormdb.DB.AutoMigrate(&models.Usuario{})
	*/
}
