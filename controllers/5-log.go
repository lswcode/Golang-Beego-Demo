package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type LogController struct {
	beego.Controller
}

func (l *LogController) Get() {
	logs.Emergency("这是Emergency日志")
	logs.Alert("这是Alert日志")
	logs.Critical("这是Critical日志")
	logs.Error("这是Error日志")
	logs.Warn("这是Warn日志")
	logs.Notice("这是Notice日志")
	logs.Info("这是Info日志")
	logs.Debug("这是Debug日志")
	l.TplName = "log/index.html"
}
