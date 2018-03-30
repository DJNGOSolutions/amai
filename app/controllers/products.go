package controllers

import (
	"amai/app/models"
	"amai/app/routes"
	"fmt"

	"github.com/revel/revel"
)

type Products struct {
	*revel.Controller
}

func (c Products) Show(code string) revel.Result {
	var product models.Product
	Gdb.Raw("SELECT code,price FROM products WHERE code = ?;", code).Scan(&product)
	//return c.RenderJSON(product) //debug
	return c.Render(product)
}

func (c Products) Delete() revel.Result {
	var products []models.Product
	Gdb.Raw("SELECT product_id,code,price FROM products").Scan(&products)
	//return c.RenderJSON(products) //debugging
	return c.Render(products)
}

func (c Products) Pop(id uint) revel.Result {
	var product models.Product
	Gdb.Raw("DELETE FROM products WHERE product_id = ?;", id).Scan(&product)
	fmt.Println("id", id)
	return c.Redirect(routes.Products.Delete())
	//return c.RenderJSON(product) //debuging
}

func (c Products) Insert(code string, price uint) revel.Result {
	pro := models.Product{Code: code, Price: price}
	Gdb.Create(&pro)
	return c.Redirect(routes.App.Index())
}
