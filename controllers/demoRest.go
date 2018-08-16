package controllers

import (
	"github.com/astaxie/beego"
)

type DemoRestController struct {
	beego.Controller
}


// @Title rest
// @Description Get rest data
// @Success 200 {object} models.ZDTCustomer.Customer
// @router /rest [get]
func (c *DemoRestController)GetOne() {
	c.Data["json"] = "Get"
	//c.Data["Key"] = key
	c.ServeJSON()
}

// @Title rest
// @Description Get rest data
// @Success 200 {object} models.ZDTCustomer.Customer
// @router /rest [put]
func (c *DemoRestController)PutOne() {
	c.Data["json"] = "Put"
	//c.Data["Key"] = key
	c.ServeJSON()
}


// @Title rest
// @Description Get rest data
// @Success 200 {object} models.ZDTCustomer.Customer
// @router /rest [post]
func (c *DemoRestController)PostOne() {
	c.Data["json"] = "Post"
	//c.Data["Key"] = key
	c.ServeJSON()
}
