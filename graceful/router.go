package main

import (
	maincontroller "common/controller/main"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func LoadRouter() {
	// 官网
	beego.Router("/", &maincontroller.Controller{})
	beego.Router("/intro/:name", &maincontroller.Controller{}, "*:Intro")
	beego.Router("/app", &maincontroller.Controller{}, "*:App")
	beego.Router("/index.html", &maincontroller.Controller{})
	beego.Router("/about.html", &maincontroller.Controller{}, "*:About")
	beego.Router("/mail.html", &maincontroller.Controller{}, "*:Mail")
	beego.Router("/products.html", &maincontroller.Controller{}, "*:Products")
	beego.Get("/help", func(ctx *context.Context) { // 修复官家端 使用帮助 404
		ctx.Redirect(302, "/")
	})

	// aliyu 健康检测
	beego.Router("/health", &maincontroller.Controller{}, "*:Check")

	// 静态文件
	beego.SetStaticPath("/static", "../static")

}
