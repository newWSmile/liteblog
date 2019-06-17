package routers

import (
	"github.com/astaxie/beego"
	"liteblog/controllers"
)

func init() {
	//注解路由 需要调用Include。
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(&controllers.IndexController{})
}
