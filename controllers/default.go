package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func init()  {

}

type MainController struct {
	BaseController
}

func (c *MainController) Index() {
	c.TplName = "index.html"
}

func (c *MainController) Login() {
	logs.Info("get string username:",c.GetString("username"))
	c.SetUsername(c.Ctx.ResponseWriter, c.Ctx.Request,c.GetString("username"))
	c.Data["WsPort"] = beego.AppConfig.String("httpport")
	c.TplName = "live.html"
}

func (c *MainController) Camera()  {
	c.Data["WsPort"] = beego.AppConfig.String("httpport")
	c.TplName = "camera.html"
}
