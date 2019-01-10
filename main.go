package main

import (
	"flag"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"scaffold_go/conf"
	"scaffold_go/database"
	"scaffold_go/log"
)
type Product struct {
	gorm.Model
	Code string
	Price int
}

var param_test1 string

/**
获得命令行传递的参数
 */
func init(){
	flag.StringVar(&param_test1, "test1", "none", "help message for test1")
	flag.Parse()
}

func main() {
	// 初始化配置信息
	conf.Initial("./conf.ini")
	// log 的输出
	// TODO: 目前还有问题，没有办法使用全局变量。每一次实例化都会生成一个对象
	logger := log.New()
	logger.Info("1111111")

	// 命令行参数的获得
	logger.Info("param: " + param_test1)
	conf.AddParam(param_test1)
	logger.Info(conf.GetConf().Param_test1)


	instance := database.GetInstence()
	conPool := instance.InitDbPool()
	if conPool == false{
		panic("connect error")
	}
	gdb := instance.GetMysqlDB()
	gdb.AutoMigrate(&database.Email{})

	}
