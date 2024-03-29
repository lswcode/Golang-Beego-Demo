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

	res := map[string]interface{}{ // 返回键值对格式给前端，map或结构体都可以，推荐使用map，因为结构体的字段名一般都是大写，需要使用tag标签进行转换
		"code": 200, "msg": "添加成功",
	}

	a.Data["json"] = res // 表示以json格式发送数据给前端，会设置响应体头中的content-type为application/json
	// a.Data["xml"]  xml格式，会设置content-type为application/xml
	a.ServeJSON() // map格式必须转换成JSON格式才能发送
}
