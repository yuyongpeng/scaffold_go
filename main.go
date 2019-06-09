package main

import (
	"flag"
	cmd "scaffold_go/cmd"
	"scaffold_go/conf"
	"scaffold_go/database"
	"scaffold_go/log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)

type Product struct {
	gorm.Model
	Code  string
	Price int
}

var param_test1 string

/**
获得命令行传递的参数
*/
func init() {
	flag.StringVar(&param_test1, "test1", "none", "help message for test1")
	flag.Parse()
}

//只是模拟一个错误
func openFile() ([]byte, error) {
	//return nil, &err.ValueError{"文件错误，自定义"}
	return nil, errors.New("文件错误，自定义")
}
func main() {
	// 命令行的参数处理
	cmd.Execute()

	// 初始化配置信息
	conf.Initial("./conf.ini")
	// log 的输出
	// TODO: 目前还有问题，没有办法使用全局变量。每一次实例化都会生成一个对象
	logger := log.New()
	logger.Info("1111111")

	// logger.Info(err3.Msg[1001])

	// 命令行参数的获得
	logger.Info("param: " + param_test1)
	conf.AddParam(param_test1)
	logger.Info(conf.GetConf().Param_test1)

	err := errors.New("whoops")
	//fmt.Printf("%+v", err)
	logger.Error(err)
	logger.Errorf("%+v", err)
	logger.Error("aaaaa\nbbbbbb")

	_, err2 := openFile()
	if err2 != nil {
		logger.Errorf("%+v", err2)
	}

	instance := database.GetInstence()
	conPool := instance.InitDbPool()
	if conPool == false {
		panic("connect error")
	}
	gdb := instance.GetMysqlDB()
	gdb.AutoMigrate(&database.Email{})

}
