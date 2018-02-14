package controllers

import (
	"techblog/utils"

	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Get() {
	c.Data["topHeader"] = utils.TopHeader
	// c.Layout = "templates/basic-layout.tpl"
	// c.LayoutSections = make(map[string]string)
	// c.LayoutSections["header"] = "templates/header.tpl"
	// c.LayoutSections["content"] = "main/homeView.tpl"
	// c.LayoutSections["footer"] = "templates/footer.tpl"
	c.TplName = "main/homeView.tpl"
}
