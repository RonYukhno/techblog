package controllers

import (
	"github.com/astaxie/beego"
)

var topHeader []string

type HomeController struct {
	beego.Controller
}

func DrawTemplate(m *HomeController, content string) {
	m.Layout = "templates/basic-layout.tpl"
	m.LayoutSections = make(map[string]string)
	m.LayoutSections["header"] = "templates/header.tpl"
	m.LayoutSections["content"] = content
	m.LayoutSections["footer"] = "templates/footer.tpl"
	m.TplName = content
}

func (m *HomeController) Get() {
	topHeader := []string{"Новости", "Обзоры", "Лента"}
	m.Data["topHeader"] = topHeader
	DrawTemplate(m, "main/homeView.tpl")
}
