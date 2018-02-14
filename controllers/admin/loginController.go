package adminControllers

import (
	"strconv"
	"techblog/models"
	"techblog/utils"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["topHeader"] = utils.TopHeader
	c.TplName = "admin/login.tpl"
}

func (c *LoginController) Post() {
	var data models.Data
	c.ParseForm(&data)

	// Check if user is logged in
	/*session := c.StartSession()
	userID := session.Get("UserID")

	if userID != nil {
		// User is logged in already, display another page
		return
	}*/

	var isLogin = data.CheckPaswrd()

	c.Ctx.Output.Body([]byte("isLogin: " + strconv.FormatBool(isLogin)))
}
