package controllers

import (
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
)

type Admin struct {
	gormc.Controller
}

/*
func (c Admin) CreateSubject(name string) revel.Result {

	subject := models.Subject{SubjectName: name}
	c.DB.Create(&subject)

	return c.RenderJSON(subject)
}

func (c Admin) CreateCategory(name string) revel.Result {
	topic := models.Topic{TopicName: name}
	c.DB.Create(&topic)

	return c.RenderJSON(topic)
}

func (c Admin) CreateAcademicLevel(name string) revel.Result {
	level := models.AcademicLevel_User{AcademicLevel: name}
	c.DB.Create(&level)
	return c.RenderJSON(level)
}

func (c Admin) Pop(id uint) revel.Result {
	//var user models.User
	//c.DB.Raw(`DELETE FROM public."user" WHERE id = ?;`, id)
	//return c.RenderJSON(c.DB.Raw(`DELETE FROM public."user" WHERE id = ?;`, id))
	return c.RenderJSON(c.DB.Unscoped().Delete(models.User{}, "id = ?", id))
	//return c.Redirect("/crud")
	//return c.RenderJSON(product) //debuging
}*/
