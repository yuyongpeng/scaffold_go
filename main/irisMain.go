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
Description:
*/
package main

import (
	"crypto/md5"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"io"
	"os"
	"scaffold_go/config"
	sv "scaffold_go/service"
	"scaffold_go/utils"
	"scaffold_go/web/views/upload"
	"strconv"
	"strings"
	"time"
)

const maxSize = 5 << 20 // 5MB

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
	//将upload_form.html提供给客户端。
	app.Get("/upload", func(ctx iris.Context) {
		//创建一个令牌（可选）。
		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		//使用令牌渲染表单以供您使用。
		ctx.ViewData("", token)
		//或者在`View`方法中添加第二个参数。

		//令牌将作为{{.}}传递到模板中。

		//ctx.View("/upload/upload_form.html", token)
		//ctx.ViewData(Body, token)
		//web.Body
		ctx.HTML(upload.Body)
	})

	//处理来自upload_form.html的请求数据处理
	app.Post("/upload", iris.LimitRequestBodySize(maxSize+1<<20), func(ctx iris.Context) {
		// Get the file from the request.
		file, _, err := ctx.FormFile("uploadfile")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}
		defer file.Close()
		//fname := info.Filename
		//创建一个具有相同名称的文件
		//假设你有一个名为'uploads'的文件夹
		randomStr := utils.Krand(16, utils.KC_RAND_KIND_ALL)
		filePath := config.Cfg.Iris.Upload + string(randomStr)

		out, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}
		defer out.Close()
		//
		io.Copy(out, file)

		// 解析CSV文件，导入到数据库中
		errorLine, success := sv.CsvService(filePath)

		var errBody string = ""
		for _, line := range errorLine {
			errBody = errBody + strings.Join([]string{"<br>", line, "</br>"}, "")
		}
		os.Remove(filePath)
		if len(errorLine) == 0 {
			header := "<h1>全部导入完毕</h1>"
			msg1 := "<h3>导入成功：" + strconv.Itoa(success) + "</h3>"
			msg2 := "<h3>导入失败：" + strconv.Itoa(len(errorLine)) + "</h3>"
			body := header + msg1 + msg2
			ctx.HTML(body)
		} else {
			header := "<h1>导入数据结果</h1>"
			msg1 := "<h3>导入成功：" + strconv.Itoa(success) + "</h3>"
			msg2 := "<h3>导入失败：" + strconv.Itoa(len(errorLine)) + "</h3>"
			errorMsg := "<h3>导入出错的数据：</h3>"
			body := header + msg1 + msg2 + errorMsg + errBody
			ctx.HTML(body)
		}
	})
	//在http//localhost:8080启动服务器，上传限制为5MB。
	app.Run(iris.Addr(":8085"), iris.WithPostMaxMemory(maxSize))
}
