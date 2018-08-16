package controllers

import (
	"github.com/astaxie/beego"
)

type DemoController struct {
	beego.Controller
}


// @Title getContent
// @Description Get获取数据
// @Success 200 {object} models.ZDTCustomer.Customer
// @router /getContent [get]
func (c *DemoController)GetContent() {
	c.Data["json"] = "111111"
	//c.Data["Key"] = key
	c.ServeJSON()
}
