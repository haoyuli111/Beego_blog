package routers

import (
	"github.com/astaxie/beego"
	"lcc/blog/controllers"
)

func init() {
	beego.Include(&controllers.UserController{})
	//添加子路由
	beego.AddNamespace(
		beego.NewNamespace(
			"note",
			beego.NSInclude(&controllers.NoteController{}),
		),
	)
}
