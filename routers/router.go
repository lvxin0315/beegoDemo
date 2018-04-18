package routers

import (
	"beegoDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{},"get:Index")
	beego.Router("/login", &controllers.MainController{},"post:Login")
	beego.Router("/camera", &controllers.MainController{},"get:Camera")
	beego.Router("/chat", &controllers.WSController{})
}
