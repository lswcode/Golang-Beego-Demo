package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

func (f *FileController) Get() {
	f.TplName = "file/index.html"
}

func (f *FileController) Post() {
	// fmt.Println("获取请求信息", f.Ctx.Request.Header)
	// 获取上传的文件
	fo, h, err := f.GetFile("file") // 参数写前端input中的name属性的值，返回文件对象，文件头，err; 文件头中可以获取文件名
	if err != nil {
		fmt.Println("f.GetFile报错:", err)
		return
	}
	defer fo.Close() // 使用完成后释放文件资源

	// 使用时间戳命名，防止文件名重复
	timeUnix := time.Now().Unix()              // 时间戳是int64类型
	timeStr := strconv.FormatInt(timeUnix, 10) // 将int64类型转为字符串，因为数字不能直接和字符串拼接

	// 保存获取到的文件
	err2 := f.SaveToFile("file", "storage/"+timeStr+"-"+h.Filename) // 设置保存的路径和文件名，第一个参数是前端input中的name属性的值
	if err2 != nil {
		fmt.Println("SaveToFile保存文件失败:", err2)
	}
	f.Data["json"] = map[string]interface{}{"code": 200, "msg": "上传成功"} // Data是一个map类型数据，用来存储要输出给前端的数据，[ ]内的键表示值的类型
	f.ServeJSON()
}
