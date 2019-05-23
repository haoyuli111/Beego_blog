package main

import (
	"encoding/gob"
	"lcc/blog/models"
	_ "lcc/blog/routers"
	"github.com/astaxie/beego"
)

func main() {
	//开启session的功能
	initSession()
	beego.Run()
}

func initSession()  {
	//用户gob编码
	gob.Register(models.User{})
	//开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	//更改session配置-----定义session的名字
	beego.BConfig.WebConfig.Session.SessionName="lcc"
	//更改session的保存方式以及目录----官网有
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig="data/session"
}

