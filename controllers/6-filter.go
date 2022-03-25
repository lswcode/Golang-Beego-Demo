package controllers

import "github.com/astaxie/beego/context"

func FilterFun(ctx *context.Context) {
	// 获取用户访问的URL
	if ctx.Request.RequestURI != "/test" {
		ctx.Redirect(302, "/test")
		ctx.WriteString("页面重定向到test")
	}

}
