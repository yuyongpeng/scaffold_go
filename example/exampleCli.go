/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/scaffold_go
Date:         2019-06-09 10:38:18
LastEditors:
LastEditTime: 2019-06-09 10:38:18
Description:
*/
package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	var language string
	var properties string
	var name string

	app := cli.NewApp()
	app.Name = "test"      // 程序名称
	app.Usage = "cli 测试程序" // 程序的描述
	app.Version = "1.0.2"  //程序的版本号
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "yuyongpeng",
			Email: "yuyongpeng@hotmail.com",
		},
	}
	app.Copyright = "(c) 1999 Serious Enterprise"
	// 全局的flag参数
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port, p",
			Value:  8000,
			Usage:  "listening port",
			EnvVar: "PORT", // 从环境变量中读取数据
		},
		cli.StringFlag{
			Name:        "lang, l",
			Value:       "english",
			Usage:       "read from 'File'",
			Destination: &language, // 把获得的变量写入language 变量
		},
	}

	// 使用到的命令和子命令
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to the list",
			// 给add命令添加flag
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "properties, p",
					Usage:       "属性文件",
					Destination: &properties,
				},
			},
			// add命令执行的动作，可以获得add命令里面的flag和
			Action: func(c *cli.Context) error {
				k := c.Int("port")
				fmt.Println("", k)
				//获得 properties的值
				p := c.String("properties")
				fmt.Print("", p)
				//获得 add命令的 arguments
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "template",
			Aliases: []string{"t"},
			Usage:   "options for task templates",
			//在 template 命令下面添加 add 和 remove 子命令
			Subcommands: []cli.Command{
				{
					Name:  "add",
					Usage: "add a new template",
					// 给add命令添加flag
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:        "name, n",
							Usage:       "姓名",
							Destination: &name,
						},
					},
					//add 子命令对应的处理动作
					Action: func(c *cli.Context) error {
						nm := c.String("name")
						fmt.Print("", nm)
						fmt.Println("new task template: ", c.Args().First())
						return nil
					},
				},
				{
					Name:  "remove",
					Usage: "remove an existing template",
					Action: func(c *cli.Context) error {
						fmt.Println("removed task template: ", c.Args().First())
						return nil
					},
				},
			},
		},
	}
	// 在命令之前运行
	app.Before = func(c *cli.Context) error {
		fmt.Print("app before")
		return nil
	}
	// 在命令之后运行
	app.After = func(c *cli.Context) error {
		fmt.Print("app after")
		return nil
	}
	//对 Falgs的参数进行排序输出
	sort.Sort(cli.FlagsByName(app.Flags))
	//对command的参数排序
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
