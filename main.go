package main

import (
	"beego_test/controllers"
	_ "beego_test/routers" // 导入项目内部的包时，使用项目名开头

	_ "beego_test/db" // 启动数据库

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

func main() {
	// beego.SetStaticPath("/static", "static")  // 设置静态资源路径，第一个参数是访问的路径，第二个参数是项目目录中的真实路径，系统默认设置了这一行代码
	// beego.SetStaticPath("/", "views") // 设置views目录为静态资源目录，localhost:8080/css/chunk-11css请求，会去views中查找资源，此时就可以把css，js等文件夹放在views中了

	// ----------------------------------------------------------------------------------------------------------------
	// 过滤器:对特定的请求进行过滤限制，第一个参数: 过滤的目标路由规则，第二个参数: 过滤器触发的位置，第三个参数: 处理函数
	beego.InsertFilter("/cs/*", beego.BeforeRouter, controllers.FilterFun) // 过滤器写在main中的beego.Run之前
	// 处理函数的参数必须是 (ctx *context.Context)
	// -------------------------------------------------------------

	orm.RunCommand() // 运行orm，数据库使用orm的前提
	// logs.SetLogger(logs.AdapterConsole) // 启用logs日志系统，console等价于logs.AdapterConsole，表示输出到控制台
	// AdapterFile      = "file"
	// logs.SetLogger(logs.AdapterFile, `{"filename":"logs/beego_project.log"}`) // 将所有日志写入一个文件
	// maxlines 每个文件保存的最大行数，默认值 1000000
	// maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB 2^28 1<< 3 2^3
	// daily 是否按照每天 logrotate，默认是 true
	// maxdays 文件最多保存多少天，默认保存 7 天
	// rotate 是否开启 logrotate，默认是 true
	// level 日志保存的时候的级别，默认是 Trace 级别
	// perm 日志文件权限 4(读权限)2(写权限)1(执行权限)
	logs.SetLogger(logs.AdapterMultiFile, // 日志分类别，多文件写入，各个级别不同的日志生成单独的文件
		`{"filename":"logs/test.log","separate":
		 ["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)
	// ---------------------------------------
	beego.Run() // 其它的配置都要在beego.Ru()之前就配置好，beego.Run()就表示服务已经开始运行，在此之后的配置就无法生效了
}
