package controllers

import (
	"github.com/pdmp/amai/app/models"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type Session struct {
	gormc.Controller
}

func (c Session) getSessionByPlace(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT Session_Date, Session_Time_Start, Session_Time_End, SessionFee FROM Session WHERE Id_Session_Place = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionTopicBySubject(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT Session_Date, Session_Time_Start, Session_Time_End, Session_Fee FROM Session WHERE Id_Subject = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionByTopic(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT Session_Date, Session_Time_Start, Session_Time_End, Session_Fee FROM Session WHERE Id_Topic = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionByType(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT Session_Date, Session_Time_Start, Session_Time_End, Session_Fee FROM Session WHERE Id_Type_Session = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}

func (c Session) getSessionByPayment(id uint) revel.Result {
	var session models.Session
	c.DB.Raw("SELECT Session_Date, Session_Time_Start, Session_Time_End, Session_Fee FROM Session WHERE Id_Type_Payment = ?;", id).Scan(&session)
	return c.RenderJSON(c.Response.Status)
}
