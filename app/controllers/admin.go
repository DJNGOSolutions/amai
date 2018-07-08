package controllers

import (
	"github.com/pdmp/amai/app/models"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type Admin struct {
	gormc.Controller
}

func (c Admin) Subjects() revel.Result {
	var subjects []*models.Subject

	c.DB.Find(&subjects)

	return c.RenderJSON(subjects)
}

func (c Admin) CreateSubject(name string) revel.Result {

	subject := models.Subject{SubjectName: name}
	c.DB.Create(&subject)

	return c.RenderJSON(subject)
}

func (c Admin) Category() revel.Result {
	var topics []*models.Topic

	c.DB.Find(&topics)

	return c.RenderJSON(topics)
}

func (c Admin) CreateCategory(name string) revel.Result {
	topic := models.Topic{TopicName: name}
	c.DB.Create(&topic)

	return c.RenderJSON(topic)
}

func (c Admin) AcademicLevel() revel.Result {
	var levels []*models.AcademicLevel_User

	c.DB.Find(&levels)

	return c.RenderJSON(levels)
}

func (c Admin) CreateAcademicLevel(name string) revel.Result {
	level := models.AcademicLevel_User{AcademicLevel: name}
	c.DB.Create(&level)
	return c.RenderJSON(level)
}

func (c Admin) Gender() revel.Result {
	var genders []*models.Gender_User

	c.DB.Find(&genders)

	return c.RenderJSON(genders)
}
