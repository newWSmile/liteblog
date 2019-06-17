package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

// 约定：如果子controller 存在NestPrepare()方法，就实现了该接口，
type NestPreparer interface {
	NestPreparer()
}

type BaseController struct {
	beego.Controller
}

func (ctx *BaseController) Prepare() {
	log.Println("BaseController")

	//将页面路径 保存到 Path变量里面
	ctx.Data["Path"] = ctx.Ctx.Request.RequestURI

	// 判断子类是否实现了NestPreparer接口，如果实现了就调用接口方法。
	if app, ok := ctx.AppController.(NestPreparer); ok {
		app.NestPreparer()
	}

}
