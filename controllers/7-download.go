package controllers

import (
	"github.com/astaxie/beego"
)

type DownloadController struct {
	beego.Controller
}

func (this *DownloadController) Get() { // 可以使用this来指代控制器对象，统一命名为this，看起来比较清晰
	this.Ctx.Output.Download("storage/下载文件.txt", "下载成功.txt") // 第一个参数是文件路径，第二个参数是下载后显示的文件名
}
