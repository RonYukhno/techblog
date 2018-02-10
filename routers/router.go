package routers

import (
	"github.com/astaxie/beego"
	"techblog/controllers/main"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
}
