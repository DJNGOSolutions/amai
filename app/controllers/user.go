package controllers

import (
	"fmt"

	"github.com/pdmp/amai/app/models"

	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
)

type User struct {
	gormc.Controller
}

func (c User) Show() revel.Result {
	var users []*models.User
	//var prods models.Product
	//Gdb.Select("code", "price").Find(&prods).Scan(&users)
	c.DB.Raw("SELECT * FROM public.user;").Scan(&users)
	return c.RenderJSON(users)
}

func (c User) Crud() revel.Result {
	return c.RenderFileName("bento/dist/index.html", revel.NoDisposition)
}

func (c User) GetUser(code string) revel.Result {
	var user models.User
	c.DB.Raw("SELECT * FROM public.user WHERE id = ?;", code).Scan(&user)
	//return c.RenderJSON(product) //debug
	return c.RenderJSON(user)
}

func (c User) Delete() revel.Result {
	var users []models.User
	c.DB.Raw("SELECT UserName FROM public.user").Scan(&users)
	//return c.RenderJSON(users) //debugging
	return c.Render(users)
}

func (c User) Pop(id uint) revel.Result {
	//var user models.User
	c.DB.Raw("DELETE FROM public.user WHERE id = ?;", id)
	fmt.Println("id", id)
	return nil
	//return c.Redirect("/crud")
	//return c.RenderJSON(product) //debuging
}

func (c User) Insert( /*code string, price uint */ user models.User) revel.Result {
	// pro := models.Product{Code: code, Price: price}
	c.DB.Create(&user)
	log := c.Log.New("insert", 1)
	//log.Debug("Inserting:", user.UserName, user.UserAge, user.UserEmail, nil)
	log.Debug("Inserting:", user)

	//return nil
	return c.RenderJSON(c.Response.Status)
	//return c.Redirect(routes.App.Index())
}

func (c User) getUserByGender(id uint) revel.Result {
	var user models.User
	c.DB.Raw("SELECT User_Photo, User_Name, User_Age, User_Email, User_Description FROM public.User WHERE Id_Gender_User = ?;", id).Scan(&user)
	return c.RenderJSON(c.Response.Status)
}

func (c User) getUserAcademicLevel(id uint) revel.Result {
	var user models.User
	c.DB.Raw("SELECT User_Photo, User_Name, User_Age, User_Email, User_Description FROM public.User WHERE Id_Academic_Level_User = ?;", id).Scan(&user)
	return c.RenderJSON(c.Response.Status)
}

func (c User) getUserByRole(id uint) revel.Result {
	var user models.User
	c.DB.Raw("SELECT User_Photo, User_Name, User_Age, User_Email, User_Description FROM public.User WHERE Id_Role_User = ?;", id).Scan(&user)
	return c.RenderJSON(c.Response.Status)
}

func (c User) getUserByRate(id uint) revel.Result {
	var user models.User
	c.DB.Raw("SELECT User_Photo, User_Name, User_Age, User_Email, User_Description FROM public.User WHERE Id_Rate_User = ?;", id).Scan(&user)
	return c.RenderJSON(c.Response.Status)
}
