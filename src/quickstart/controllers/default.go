package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "wangwentao@le.com"
	c.TplName = "index.html"
	//c.TplName = "index.tpl"
}
