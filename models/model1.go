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

func init() {

	orm.RegisterModel(new(User)) // 注册模型
}
