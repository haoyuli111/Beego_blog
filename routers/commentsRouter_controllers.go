package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["lcc/blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["lcc/blog/controllers:NoteController"],
		beego.ControllerComments{
			Method:           "Edit",
			Router:           `/edit`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["lcc/blog/controllers:NoteController"] = append(beego.GlobalControllerRouter["lcc/blog/controllers:NoteController"],
		beego.ControllerComments{
			Method:           "Save",
			Router:           `/save`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["lcc/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["lcc/blog/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Del",
			Router:           `/del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["lcc/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["lcc/blog/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/get`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["lcc/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["lcc/blog/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetDetails",
			Router:           `/getdetails`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["lcc/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["lcc/blog/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["lcc/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["lcc/blog/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Logout",
			Router:           `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["lcc/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["lcc/blog/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Register",
			Router:           `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["lcc/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["lcc/blog/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UpdatedInfo",
			Router:           `/updatedInfo`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["lcc/blog/controllers:UserController"] = append(beego.GlobalControllerRouter["lcc/blog/controllers:UserController"],
		beego.ControllerComments{
			Method:           "UpdatedPwd",
			Router:           `/updatedPwd`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

}
