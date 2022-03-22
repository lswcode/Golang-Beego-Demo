package db

import (
	_ "beego_test/models" // 导入项目内部的包时，使用项目名开头
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 数据库的使用
	// 第一步
	username := beego.AppConfig.String("username") // 获取app.conf中的数据
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	database := beego.AppConfig.String("database")

	// -------------------------------------------------------
	// 第二步
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", username, password, host, port, database)
	orm.RegisterDriver("mysql", orm.DRMySQL)                            // 第一个参数是数据库名称，第二个参数是数据库类型，使用orm包中的mysql类型
	err := orm.RegisterDataBase("default", "mysql", datasource, 30, 30) // 开始连接数据库
	// 第一个参数是自定义的数据库的别名，第二个参数是上一行orm.RegisterDriver中的第一个参数，第三个参数是连接数据库的字符串
	// ORM必须注册一个别名为default的数据库，作为默认使用
	// 第四个和第五个参数可以省略，设置最大空闲连接和最大数据库连接
	if err != nil {
		fmt.Println("数据库连接报错:", err)
	}

	// 第三步创建数据模型models
	// -------------------------------------------------------
	// beego的orm运行参数
	// go run beego_rom.go orm sqlall   显示orm创建的表时对应的sql语句
	name := "default"                           // 操作的数据库的别名
	force := false                              // 强制创建表，如果表重名，先删除旧的再创建表，一般设置为false
	verbose := true                             // 打印执行的过程，即自动执行 go run beego_rom.go orm sqlall
	err2 := orm.RunSyncdb(name, force, verbose) // 设置orm运行参数

	if err2 != nil {
		panic(err)
	}

	// orm.Debug = true

}
