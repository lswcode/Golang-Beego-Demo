package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller // 所有的请求响应处理，都继承自beego.Controller这个内置的结构体
}

// 同一个控制器中，不能有相同的请求方法，因为处理函数都是没有名字的，同一控制器下，通过请求方式来区分同一路由的不同处理函数
func (u *UserController) Get() { // 如果有两个Get方法，那beego就不知道找哪个进行处理了
	fmt.Println("获取请求信息", u.Ctx.Request.RequestURI)
	id := u.GetString("id") // 第一种 获取路径参数，参数名要和前端约好
	fmt.Println("获取到的路径参数id:", id)

	id2, err := u.GetInt("id") // 第二种 限制路径参数的格式，如果格式不符，则返回限制的格式的零值
	if err != nil {
		fmt.Println("GetInt报错:", err)
	}
	fmt.Println("限制参数为int:", id2)

	u.TplName = "user/user.html"
}

type Student struct {
	// go get github.com/astaxie/beego/validation  可以使用这个包进行表单格式验证
	// 反射解析结构体标签tag
	Username string `form:"username"` // 这里字段名必须大写，才能被beego包访问，下面才能使用beego包中的工具函数将数据直接保存数据到结构体中
	Age      string `form:"age"`      // 字段名大写后，和form表单中传递的属性名就不同了，需要使用结构体的tag标签` `进行处理
}

func (u *UserController) Post() {
	username := u.GetString("username") // 获取post表单提交的数据
	age := u.GetString("age")
	fmt.Println("username:", username)
	fmt.Println("age:", age)
	fmt.Printf("age类型%T\n", age)

	// --------------------------------------------
	// 解析表单数据，一次性全部保存到结构体中
	fmt.Println("获取请求信息", u.Ctx.Request.RequestURI)
	// u.Ctx.Request 所有的请求信息
	// u.Ctx.Request.Header 请求头
	// u.Ctx.Request.Host 请求的主机
	// u.Ctx.Request.Method 请求的方法
	newStudent := Student{}
	err := u.ParseForm(&newStudent) // 注意
	if err != nil {
		fmt.Println("ParseForm报错:", err)
	}
	fmt.Println(newStudent)
	fmt.Println("保存到结构体中的name和age", newStudent.Username, newStudent.Age)

	u.Ctx.WriteString("提交成功")
}
