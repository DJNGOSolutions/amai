package controllers

import (
	"github.com/pdmp/amai/app/models"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type Session struct {
	gormc.Controller
}

func (c Session) getSessionById(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT SessionTopic FROM Session WHERE Id = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionTopicBySubject(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT SessionTopic FROM Session WHERE IdSubject = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionTopicById(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT SessionTopic FROM Session WHERE IdTopic = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionTopicByUserId(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT SessionTopic FROM Session WHERE IdUser = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionTopicByInstitution(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT SessionTopic FROM Session WHERE IdSubject = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}
