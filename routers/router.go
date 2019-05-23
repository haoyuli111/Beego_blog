package routers

import (
	"github.com/astaxie/beego"
	"lcc/blog/controllers"
)

func init() {
	beego.Include(&controllers.UserController{})
}
