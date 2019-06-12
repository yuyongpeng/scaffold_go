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
	"io"
	"os"
	"scaffold_go/config"
	"scaffold_go/web/views/upload"
	"strconv"
	"time"
)

const maxSize = 5 << 20 // 5MB


func main() {

	app := iris.New()
	app.RegisterView(iris.HTML(config.Cfg.Iris.Html, ".html"))
	fmt.Println(config.Cfg)
	// 初始化Iris的配置
	configuration := config.InitIrisConfiguration()
	iris.WithConfiguration(configuration)


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
		file, info, err := ctx.FormFile("uploadfile")
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}
		defer file.Close()
		fname := info.Filename
		//创建一个具有相同名称的文件
		//假设你有一个名为'uploads'的文件夹
		out, err := os.OpenFile(config.Cfg.Iris.Upload+fname,
			os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}
		defer out.Close()
		io.Copy(out, file)
	})
	//在http//localhost:8080启动服务器，上传限制为5MB。
	app.Run(iris.Addr(":8085"), iris.WithPostMaxMemory(maxSize))
}
