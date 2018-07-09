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

func (c App) Login(email string, pass string) revel.Result {
	var user models.User
	//return c.RenderJSON(c.DB.Where("user_email = ?", email).Find(&user))
	c.DB.Where("user_email = ?", email).Find(&user)
	//return c.RenderJSON(&user)
	fmt.Println("hash", user.UserPassword, "\npass -> "+pass, "\nbyte", []byte(pass))
	if user.UserEmail != "" {
		err := bcrypt.CompareHashAndPassword([]byte(user.UserPassword), []byte(pass))
		fmt.Println(err)
		if err == nil {
			if user.IdRoleUser == 1 {
				c.Session["user"] = user.UserName
				c.Session.SetDefaultExpiration()
			}
			token := encodeToken(email)
			return c.RenderJSON(token)
		}
	}
	return c.RenderFileName("bento/dist/index.html", revel.NoDisposition)
}

func (c App) Register(user models.User) revel.Result {

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(user.UserPassword), bcrypt.DefaultCost)
	user.UserPassword = string(hashPass)
	c.DB.Create(&user)

	//return nil
	return c.RenderJSON(c.Response.Status)
	//return c.Redirect(routes.App.Index())
}

func (c App) Help() revel.Result {
	return c.Render()
}

func (c App) HashTest(pass string) revel.Result {
	hashPass, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	c.DB.Model(&models.User{}).Where("id = ?", 31).Update("user_password", hashPass)
	//c.DB.Raw("UPDATE public.user SET user_password = ? WHERE id = 31", hashPass)
	return c.RenderText(string(hashPass))
}
