package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
)

func init() {
        beego.Router("/", &controllers.MainController{})
	beego.Router("/api/getInfo", &controllers.UserController{}, "get:UserGet")
	beego.Router("/api/updateUser", &controllers.UserController{}, "post:Update")
	//beego.Router("/api",&RestController{},"get,post:ApiFunc")
	//beego.Router("/simple",&SimpleController{},"get:GetFunc;post:PostFunc")
}
