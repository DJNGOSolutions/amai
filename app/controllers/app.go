package controllers

import (
	"fmt"

	"github.com/pdmp/amai/app/models"
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	gormc.Controller
}

var (
	url = "bento/dist/index.html"
)

func (c App) Index() revel.Result {
	return c.RenderFileName(url, revel.NoDisposition)
}

func (c App) Login() revel.Result {
	username := c.Params.Form.Get("username")
	pass := c.Params.Form.Get("password")
	var user models.UserT
	//return c.RenderJSON(c.DB.Where("user_username = ?", username).Find(&user))
	c.DB.Where("user_name = ?", username).Find(&user)
	//return c.RenderJSON(&user)
	fmt.Println("hash", user.UserPassword, "\npass -> "+pass, "\nbyte", []byte(pass))
	if user.UserName != "" {
		err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(pass))
		if err == nil {
			//TODO INSERT TYPE ON DATABASE
			if user.IdTipo == 1 {
				c.Session["user"] = user.UserName
				c.Session.SetDefaultExpiration()
			}
			token := encodeToken(username)
			return c.RenderJSON(token)
		}
	}
	return c.RenderJSON("auth failed")
	//return c.RenderFileName("bento/dist/index.html", revel.NoDisposition)
}
func (c App) Register(user models.UserT) revel.Result {
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
	user.UserPassword = hashPass
	c.DB.Create(&user)

	return c.RenderJSON(c.Response.Status)
	//return c.Redirect(routes.App.Index())
}

func (c App) Insert() revel.Result {
	username := c.Params.Form.Get("username")
	pass := c.Params.Form.Get("password")

	password, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	user := models.UserT{UserName: username, UserPassword: password}
	c.DB.Create(&user)

	return c.RenderJSON(user)
}
