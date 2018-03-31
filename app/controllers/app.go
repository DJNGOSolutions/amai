package controllers

import (
	"amai/app/models"
	"amai/app/routes"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	var products []*models.Product
	//var prods models.Product
	//Gdb.Select("code", "price").Find(&prods).Scan(&products)
	Gdb.Raw("SELECT * FROM products;").Scan(&products)
	return c.Render(products)
}

func (c *App) Insert() revel.Result {
	pro := models.Product{Code: "a", Price: 100}
	Gdb.Create(&pro)
	return c.Redirect(routes.App.Index())
}

func (c App) Help() revel.Result {
	return c.Render()
}
