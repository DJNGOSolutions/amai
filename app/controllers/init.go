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

	gormdb.DB.AutoMigrate(&models.Assistance{})
	gormdb.DB.AutoMigrate(&models.Session{})
	gormdb.DB.AutoMigrate(&models.Gender_User{})
	gormdb.DB.AutoMigrate(&models.Place_Session{})
	gormdb.DB.AutoMigrate(&models.Subject{})
	gormdb.DB.AutoMigrate(&models.AcademicLevel_User{})
	gormdb.DB.AutoMigrate(&models.Role_User{})
	gormdb.DB.AutoMigrate(&models.Topic{})
	gormdb.DB.AutoMigrate(&models.Type_Session{})
	gormdb.DB.AutoMigrate(&models.TypePayment_Session{})
	gormdb.DB.AutoMigrate(&models.User{})
	gormdb.DB.AutoMigrate(&models.RatexUser{})
	gormdb.DB.AutoMigrate(&models.Classroom{})
	gormdb.DB.AutoMigrate(&models.UserxClassroom{})
	gormdb.DB.AutoMigrate(&models.Subscription{})
	gormdb.DB.AutoMigrate(&models.Category{})
}
