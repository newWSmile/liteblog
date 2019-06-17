package routers

import (
	"github.com/astaxie/beego"
	"liteblog/controllers"
)

func init() {
	beego.ErrorController(&controllers.ErrorController{})
	beego.Include(
		&controllers.IndexController{},
		&controllers.UserController{},
	)
	beego.AddNamespace(
		beego.NewNamespace(
			"note",
			beego.NSInclude(&controllers.NoteController{}),
		),
		beego.NewNamespace(
			"upload",
			beego.NSInclude(&controllers.UploadController{}),
		),
		beego.NewNamespace(
			"praise",
			beego.NSInclude(&controllers.PraiseController{}),
		),
		beego.NewNamespace(
			"message",
			beego.NSInclude(&controllers.MessageController{}),
		),
	)
}
