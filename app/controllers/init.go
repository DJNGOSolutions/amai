package controllers

import (
	"github.com/pdmp/amai/app/models"
	gormdb "github.com/revel/modules/orm/gorm/app"
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(InitDB)
}

func InitDB() {
	revel.INFO.Println("Doing DB Migrations...")
	gormdb.DB.AutoMigrate(&models.Product{})
}
