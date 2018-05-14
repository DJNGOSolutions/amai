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
	var users []*models.Usuario
	//var prods models.Product
	//Gdb.Select("code", "price").Find(&prods).Scan(&users)
	c.DB.Raw("SELECT * FROM usuario;").Scan(&users)
	return c.RenderJSON(users)
}

func (c User) Crud() revel.Result {
	return c.RenderFileName("bento/dist/index.html", revel.NoDisposition)
}

func (c User) GetUser(code string) revel.Result {
	var user models.Usuario
	c.DB.Raw("SELECT * FROM usuario WHERE id = ?;", code).Scan(&user)
	//return c.RenderJSON(product) //debug
	return c.RenderJSON(user)
}

func (c User) Delete() revel.Result {
	var users []models.Usuario
	c.DB.Raw("SELECT nombre_usuario FROM product").Scan(&users)
	//return c.RenderJSON(users) //debugging
	return c.Render(users)
}

func (c User) Pop(id uint) revel.Result {
	var user models.Usuario
	c.DB.Raw("DELETE FROM usuario WHERE id = ?;", id).Scan(&user)
	fmt.Println("id", id)
	return nil
	//return c.Redirect("/crud")
	//return c.RenderJSON(product) //debuging
}

func (c User) Insert( /*code string, price uint */ user models.Usuario) revel.Result {
	// pro := models.Product{Code: code, Price: price}
	c.DB.Create(&user)
	log := c.Log.New("insert", 1)
	log.Debug("Inserting:", user.NombreUsuario, user.EdadUsuario)
	//return nil
	return c.RenderJSON(c.Response.Status)
	//return c.Redirect(routes.App.Index())
}
