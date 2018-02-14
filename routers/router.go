package routers

import (
	"techblog/controllers/admin"
	"techblog/controllers/main"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/login", &adminControllers.LoginController{})
}
