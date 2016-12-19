package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const (
	_DB_NAME = "data/beeadmin.sql"
)

type User struct {
	Id       int64
	Name     string
	Account  string
	PhoneNum string
	Created  time.Time `orm:"index"`
}

//注册数据库
func RegisterDB() {
	//注册模型
	orm.RegisterModel(new(User))
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册数据库
	orm.RegisterDataBase("default", "mysql", "root:special005@/user?charset=utf8", 10)
}

//获取所有用户资料
func GetAllUsers() ([]*User, error) {
	o := orm.NewOrm()
	users := make([]*User, 0)
	qs := o.QueryTable("user")
	_, err := qs.All(&users)
	return users, err
}
