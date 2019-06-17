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
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"scaffold_go/config"
	"scaffold_go/database"
	"scaffold_go/elastic"
	"strings"
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

	err := elastic.InsertElastic(job, "cport_person_x")
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
	if bodyJosn, err := simplejson.NewFromReader(ctx.Request().Body); err != nil {
		fmt.Println(err)
		ctx.JSON(map[string]string{"return_code": "1008", "msg": err.Error()})
	} else {
		queryStr := bodyJosn.Get("query").MustString("")
		if queryStr == "" {
			ctx.JSON(map[string]string{"return_code": "1010", "msg": err.Error()})
		}
		job_area_id := bodyJosn.Get("job_area_id").MustInt(0)
		industry_id := bodyJosn.Get("industry_id").MustInt(0)
		job_salary := bodyJosn.Get("job_salary").MustInt(0)
		job_min_education := bodyJosn.Get("job_min_education").MustInt(0)
		job_experience := bodyJosn.Get("job_experience").MustInt(0)
		job_mode := bodyJosn.Get("job_mode").MustInt(0)
		enterprise_size := bodyJosn.Get("enterprise_size").MustInt(0)
		job_status := bodyJosn.Get("job_status").MustInt(0)
		size := bodyJosn.Get("size").MustInt(10)
		from := bodyJosn.Get("from").MustInt(0)
		modify_time_start := bodyJosn.Get("modify_time_start").MustString("0")
		modify_time_end := bodyJosn.Get("modify_time_end").MustString("0")

		// 拼接 查询 query的 过滤器
		var filter []interface{}
		if job_area_id != 0 {
			filter = append(filter, map[string]interface{}{
				"term": map[string]interface{}{
					"job_area_id": job_area_id,
				}})
		}
		if industry_id != 0 {
			filter = append(filter, map[string]interface{}{
				"term": map[string]interface{}{
					"industry_id": industry_id,
				}})
		}
		if job_salary != 0 {
			filter = append(filter, map[string]interface{}{
				"term": map[string]interface{}{
					"job_salary": job_salary,
				}})
		}
		if job_min_education != 0 {
			filter = append(filter, map[string]interface{}{
				"term": map[string]interface{}{
					"job_min_education": job_min_education,
				}})
		}
		if job_experience != 0 {
			filter = append(filter, map[string]interface{}{
				"term": map[string]interface{}{
					"job_experience": job_experience,
				}})
		}
		if enterprise_size != 0 {
			filter = append(filter, map[string]interface{}{
				"term": map[string]interface{}{
					"enterprise_size": enterprise_size,
				}})
		}
		if job_mode != 0 {
			filter = append(filter, map[string]interface{}{
				"term": map[string]interface{}{
					"job_mode": job_mode,
				}})
		}
		if job_status != 0 {
			filter = append(filter, map[string]interface{}{
				"term": map[string]interface{}{
					"job_status": 0,
				}})
		}
		// date range
		if modify_time_start != "0" || modify_time_end != "0" {
			var modify_time map[string]string = map[string]string{}
			if modify_time_start != "0" {
				modify_time["gte"] = modify_time_start  	// 大于等于
			}
			if modify_time_end != "0" {
				modify_time["lte"] = modify_time_end		// 小于等于
			}
			k := map[string]interface{}{
				"range": map[string]interface{}{
					"modify_time": modify_time,
				}}
			filter = append(filter, k)
		}

		query := map[string]interface{}{
			"query": map[string]interface{}{
				"bool": map[string]interface{}{
					"must": map[string]interface{}{
						"match": map[string]interface{}{
							"job_name": queryStr,
						},
					},
					"filter": map[string]interface{}{
						"bool": map[string]interface{}{
							"must": filter,
						},
					},
				},
			},
			"size": size, // 显示应该返回的结果数量，默认是 10
			"from": from, // 显示应该跳过的初始结果数量，默认是 0
			"sort": []interface{}{
				map[string]interface{}{
					"modify_time": map[string]string{
						"order": "desc",
					},
				},
			},
			"highlight": map[string]interface{}{
				"pre_tags": []string{
					"<tag1>", "<tag2>",
				},
				"post_tags": []string{
					"<tag1>", "<tag2>",
				},
				"fields": map[string]interface{}{
					"job_name": map[string]interface{}{
						"number_of_fragments": 0,
					},
				},
			},
		}
		// 打印 query 数据 用于测试
		if jsbyte, err := json.Marshal(query); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(strings.Repeat("=",20)+ "query json" + strings.Repeat("=",20))
			fmt.Println(string(jsbyte))
			fmt.Println(strings.Repeat("=",50))
		}

		retObj, err := elastic.QueryElastic(query, "cport_person_x")
		if err != nil {
			ctx.JSON(map[string]string{"return_code": "1009", "msg": err.Error()})
		}else{
			ctx.JSON(map[string]interface{}{"return_code": "200", "msg": "", "data": retObj})
		}
	}

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
