package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (t *TestController) Test1() { // 使用自定义路由可以就自定义方法名了
	t.TplName = "user/user.html"
}

func (t *TestController) Test2() {
	t.Ctx.WriteString("Test2函数")
}
