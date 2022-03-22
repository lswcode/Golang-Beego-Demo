package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type AxiosController struct {
	beego.Controller
}

func (a *AxiosController) Get() {
	a.TplName = "axios/index.html"
}

func (a *AxiosController) Post() {
	// 获取Ajax数据
	username := a.GetString("username")
	id := a.GetString("id")
	fmt.Println("axios中的username", username)
	fmt.Println("axios中的id", id)

	res := map[string]interface{}{ // 返回一个键值对给前端
		"code": 200, "msg": "添加成功",
	}
	a.Data["json"] = res // 这一行必须加
	a.ServeJSON()        // map格式必须转换成JSON格式才能发送
}
