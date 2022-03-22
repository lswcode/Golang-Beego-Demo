package routers

import (
	"beego_test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/postForm", &controllers.UserController{})
	beego.Router("/axios", &controllers.AxiosController{})
	beego.Router("/file", &controllers.FileController{})
}
