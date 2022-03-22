package routers

import (
	"beego_test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}) // 一个固定的路由，一个控制器，根据不同的请求方式，对应控制器中不同的方法
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/postForm", &controllers.UserController{})
	beego.Router("/axios", &controllers.AxiosController{})
	beego.Router("/file", &controllers.FileController{})
	// ------------------------------------------------------------
	// 自定义路由，为不同的请求方式设置对应的处理函数
	beego.Router("/test", &controllers.TestController{}, "get:Test1;post:Test2") // 使用分号隔开

}
