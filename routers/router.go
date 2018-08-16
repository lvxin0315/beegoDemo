// @APIVersion 1.0.0
// @Title mobile API
// @Description we are family
// @Contact lvxin0315@gmail.com

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



	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/demo",
			beego.NSInclude(
				&controllers.DemoController{},
			),
			beego.NSInclude(
				&controllers.DemoRestController{},
			),
		),
	)

	beego.AddNamespace(ns)
}
