package controllers

import (
	gormc "github.com/revel/modules/orm/gorm/app/controllers"
	"github.com/revel/revel"
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

func (c App) Help() revel.Result {
	return c.Render()
}
