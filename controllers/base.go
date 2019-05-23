package controllers

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/satori/go.uuid"
	"lcc/blog/models"
	"ticket/syserror"
)
//定义保存用户session的常量
const SESSION_USER_KEY  = "SESSION_USER_KEY"
type BaseController struct {
	beego.Controller
	//用来保存登录的用户
	User models.User
	//用户是否登录
	IsLogin bool
}

func (this *BaseController) Prepare(){
	//拿到当前访问的路径，比较是否为当前页面（用于添加下划线）
	this.Data["Path"] = this.Ctx.Request.RequestURI
	u, ok := this.GetSession(SESSION_USER_KEY).(models.User)
	//进行类型判断 如果用户是能得到的
	if ok {
		//将用户传给User
		this.User = u
		//当前处于登录状态
		this.IsLogin = true
		//将数据传输给Data（前端）
		this.Data["User"] = this.User
	}

}


//自定义一个错误
func (this *BaseController)Abort500(err error){
	this.Data["error"]=err
	this.Abort("500")
}

//确保输入的值不为空
func (this *BaseController) GetMustString(key,msg string) string {
	s :=this.GetString(key,"")
	if len(s)==0 {
		this.Abort500(errors.New(msg))
	}
	return s
}

//确保用户登录
func (this *BaseController) MustLogin() {
	if !this.IsLogin {
		this.Abort500(syserror.NoUserError{})
	}
}

//定义返回的数据类型
type H map[string] interface{}

func (this *BaseController) JsonOk(msg string) {
	this.Data["json"] = H{
		"code":0,
		"msg":msg,
	}
	this.ServeJSON()
}

func (this *BaseController) JsonOkH(msg string,data H) {
	if data==nil {
		data = H{}
	}
	data["code"] = 0
	data["msg"] = msg
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *BaseController) JsonOkData(msg string,data H) {
	this.Data["json"] = H{
		"code":0,
		"msg":msg,
		"data":data,
	}
	this.ServeJSON()
}


func (this *BaseController) UUID() string {
	u,err:=uuid.NewV4()
	if err!=nil {
		this.Abort500(syserror.New("系统错误",err))
	}
	return u.String()
}
