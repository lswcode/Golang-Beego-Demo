package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int
	Name string
	Age  int
	Addr string
}

func (u *User) TableName() string { // 给结构体添加方法，来自定义表名
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(User)) // 注册模型
}
