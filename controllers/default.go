package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"   // 这些是项目刚创建时的静态html模板中的数据，直接删除即可
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.html"
}
