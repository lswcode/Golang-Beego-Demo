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
	beego.Router("/log", &controllers.LogController{})
	beego.Router("/download", &controllers.DownloadController{}) // 对应的html页面在views文件夹中，自己使用vscode的提供的服务器运行
	// ------------------------------------------------------------
	// 自定义路由的处理函数，可以自定义控制器中的函数名，默认控制器中的函数名就是请求方法，根据请求方法找到对应的函数
	beego.Router("/test", &controllers.TestController{}, "get:Test1;post:Test2") // 使用分号隔开

}
