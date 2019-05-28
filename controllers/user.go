package controllers

import (
	"errors"
	"fmt"
	"lcc/blog/models"
	"lcc/blog/syserror"
	"strings"
)

type UserController struct {
	BaseController
}

//@router /login [post]
func (this *UserController) Login() {
	//name----必须填写，需要在base中加一个必须填写的方法以保证登录名不为空
	name := this.GetMustString("name", "用户名不能为空！")
	//password----必须填写，需要在base中加一个必须填写的方法以保证登录名不为空
	password := this.GetMustString("password", "密码不能为空！")

	user, err := models.QueryUserByNameAndPassword(name, password)
	fmt.Print(err)
	if err != nil {
		this.Abort500(syserror.New("登录失败！", err))
	}
	this.SetSession(SESSION_USER_KEY, user)
	this.JsonOk("登陆成功")
}

//@router /register [post]
func (this *UserController) Register() {
	//得到传进来的用户名，手机号，邮箱，密码，确认密码
	name := this.GetMustString("name", "用户名不能为空！")
	phone := this.GetMustString("phone", "手机号不能为空！")
	email := this.GetMustString("email", "邮箱不能为空！")
	password := this.GetMustString("password", "密码不能为空！")
	confirmPassword := this.GetMustString("confirmPassword", "确认密码不能为空！")
	//确认密码和密码相同
	if strings.Compare(password, confirmPassword) != 0 {
		this.Abort500(errors.New("确认密码和密码不相同"))
	}
	//查看用户名是否存在
	if u, err := models.QueryUserByName(name); err == nil && u.ID > 0 {
		this.Abort500(errors.New("用户名已存在"))
	}
	//查看邮箱是否存在
	if u, err := models.QueryUserByEmail(email); err == nil && u.ID > 0 {
		this.Abort500(errors.New("该邮箱已注册"))
	}
	//查看手机号是否存在
	if u, err := models.QueryUserByPhone(phone); err == nil && u.ID > 0 {
		this.Abort500(errors.New("该手机号已被注册"))
	}
	//符合注册的规则------进行注册
	if err := models.SaveUser(&models.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: password,
		Avatar:   "/static/images/lcc.png",
		Role:     1,
	}); err != nil {
		this.Abort500(syserror.New("用户注册失败", err))
	}
	this.JsonOk("注册成功")
}

//@router /logout [get]
func (this *UserController) Logout() {
	//只有登录的用户才能进行注销登录的操作
	this.MustLogin()
	this.DelSession(SESSION_USER_KEY)
	this.JsonOk("注销成功")
}

//@router /updatedInfo [post]
func (this *UserController) UpdatedInfo() {
	//根据传进来的用户名更改手机号，邮箱
	name := this.GetMustString("name", "用户名不能为空！")
	phone := this.GetMustString("phone", "手机号不能为空！")
	email := this.GetMustString("phone", "手机号不能为空！")
	//更新数据
	_, err := models.UpdatedInfoByName(name, phone, email)
	if err != nil {
		this.Abort500(syserror.New("更新失败！", err))
	}

	this.JsonOk("更新成功")

}

//@router /updatedPwd  [post]
func (this *UserController) UpdatedPwd() {
	//根据传进来的用户名更改手机号，邮箱
	name := this.GetMustString("name", "用户名不能为空！")
	password := this.GetMustString("password", "密码不能为空！")
	//更新数据
	_, err := models.UpdatedPwdByName(name, password)
	if err != nil {
		this.Abort500(syserror.New("更新失败！", err))
	}
	this.JsonOk("密码修改成功")

}

//@router /del [post]
func (this *UserController) Del() {
	//根据传来的用户名删除用户
	name := this.GetMustString("name", "用户名不能为空！")
	if err := models.DeleteUserByName(name); err != nil {
		this.Abort500(syserror.New("删除失败", err))
	}
	this.JsonOk("删除成功")
}

//@router /getdetails [get]
func (this *UserController) GetDetails() {
	name := this.GetMustString("name", "用户名不能为空！")
	user, err := models.QueryUserByName(name)
	if err != nil {
		this.Abort500(syserror.New("用户不存在", err))
	}

	//this.JsonOkH("查询成功",H{"name":user.Name,"password":user.Password,"id":user.ID,
	//	"phone":user.Phone,"email":user.Email,"avatar":user.Avatar})
	this.JsonOkData("查询成功", H{"name": user.Name, "password": user.Password, "id": user.ID,
		"phone": user.Phone, "email": user.Email, "avatar": user.Avatar})
}

//@router /get [get]
func (this *UserController) Get() {
	//查询数据库中的所有数据
	users, err := models.QueryAllUser()
	if err != nil {
		this.Abort500(syserror.New("用户不存在", err))
	}
	this.JsonOkData("查询成功", H{"list": users})
}
