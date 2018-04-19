package main

import (
	_ "beegoDemo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego/logs"
)

func main() {
	//设置跨域
	filter := cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	})
	//日志级别
	logs.SetLogger(logs.AdapterConsole, `{"level":3}`)

	beego.InsertFilter("*", beego.BeforeRouter, filter)
	beego.Run()
}

