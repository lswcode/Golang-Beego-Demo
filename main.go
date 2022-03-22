package main

import (
	_ "beego_test/routers"

	_ "beego_test/db" // 导入项目内部的包时，使用项目名开头

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {

	// 过滤器:对特定的请求进行过滤限制，第一个参数: 过滤的目标路由规则，第二个参数: 过滤器触发的位置，第三个参数: 处理函数
	// beego.InsertFilter("/cms/*",beego.BeforeRouter,controllers.xx) // 处理函数，一般都写在controllers包中
	// 处理函数的参数必须是 (ctx *context.Context)

	orm.RunCommand() // 运行orm，数据库使用orm的前提

	beego.Run()
}
