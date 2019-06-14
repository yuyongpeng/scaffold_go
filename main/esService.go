/**
      ___           ___           ___
     /\__\         /\__\         /\  \
    /:/ _/_       /:/  /        /::\  \
   /:/ /\  \     /:/  /        /:/\:\  \
  /:/ /::\  \   /:/  /  ___   /:/ /::\  \
 /:/_/:/\:\__\ /:/__/  /\__\ /:/_/:/\:\__\
 \:\/:/ /:/  / \:\  \ /:/  / \:\/:/  \/__/
  \::/ /:/  /   \:\  /:/  /   \::/__/
   \/_/:/  /     \:\/:/  /     \:\  \
     /:/  /       \::/  /       \:\__\
     \/__/         \/__/         \/__/
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-11 23:53:25
LastEditors:
LastEditTime: 2019-06-11 23:53:25
Description:  提供了 ES的查询和插入接口
*/
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"scaffold_go/config"
	"scaffold_go/database"
	"scaffold_go/elastic"
)

/**
插入职位的信息到ES中
*/
func insertJobHandler(ctx iris.Context) {
	//var job database.Job
	job := &database.Job{}
	if err := ctx.ReadJSON(job); err != nil {
		fmt.Print(err)
	}

	err := elastic.InsertJob(job)
	if err != nil {
		ctx.JSON(map[string]string{"return_code": "1008", "msg": err.Error()})
	} else {
		ctx.JSON(map[string]string{"return_code": "200", "msg": "", "data": ""})
	}
}

/**
查询职位的信息
*/
func queryJobHandler(ctx iris.Context) {

}

/**
插入应聘者的资料到ES中
*/
func insertPersonHandler(ctx iris.Context) {

}

/**
查询应聘者的资料
*/
func queryPersonHandler(ctx iris.Context) {

}

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML(config.Cfg.Iris.Html, ".html"))
	// 初始化Iris的配置
	//configuration := config.InitIrisConfiguration()
	//iris.WithConfiguration(configuration)
	customLogger := logger.New(logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Use(customLogger)

	es := app.Party("/es")
	es.Post("/job/insert", insertJobHandler)
	es.Post("/job/update", insertJobHandler)
	es.Post("/job/query", queryJobHandler)

	es.Post("/person/insert", insertPersonHandler)
	es.Post("/person/update", insertPersonHandler)
	es.Post("/person/query", queryPersonHandler)

	//在http//localhost:8080启动服务器，上传限制为5MB。
	app.Run(iris.Addr(":8085"))
}
