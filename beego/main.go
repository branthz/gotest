package main

import(
	    "github.com/astaxie/beego"
		_ "github.com/branthz/gotest/beego/router"
)

type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
    this.Ctx.WriteString("hello world")
}

func main() {
    beego.Router("/", &MainController{})
    beego.Run()
}
