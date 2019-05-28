package models

import (
	"github.com/jinzhu/gorm"
	//加上引擎
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var (
	db *gorm.DB
)

//数据库初始化逻辑
func init() {
	//创建错误变量
	var err error
	//创建
	db, err = gorm.Open("mysql", "root:123456@/db_lccblog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	//数据库不能关
	//defer db.Close()
	//同步表结构，用一次就得补充一次
	db.AutoMigrate(&User{}, &Note{})
	//如果数据库里面没有数据，就新增一条admin数据
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			Name:     "admin",
			Email:    "admin",
			Phone:    "admin",
			Password: "admin",
			Avatar:   "/static/images/admin.png",
			Role:     0,
		})
	}

}

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
