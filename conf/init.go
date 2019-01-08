<<<<<<< HEAD
package conf

import "C"
=======
package main

>>>>>>> 9f14ac8fb24acb5ef6109d89b707ec0be64a4e18
import (
	"github.com/go-ini/ini"
	"fmt"
	"os"
)
<<<<<<< HEAD

type g struct {
	// global
	Global_key string
	// mysql
	Mysql_username string
	Mysql_password string
	Mysql_port int
	Mysql_db string
	// redis
	Redis_username string
	Redis_password string
	Redis_port int
}
var G = &g{}

/**
将配置文件的数据写入全局变量中
*/
func Initial(confFile string){
=======
var global_key string

var mysql_username string
var mysql_password string
var mysql_port int
var mysql_db string

var redis_username string
var redis_password string
var redis_port int
/**
将配置文件的数据写入全局变量中
 */
 func Initial(confFile string){
>>>>>>> 9f14ac8fb24acb5ef6109d89b707ec0be64a4e18
	cfg,err := ini.Load(confFile)
	if err != nil{
		fmt.Println("Fail to read file :", confFile)
		os.Exit(1)
	}
<<<<<<< HEAD

	// Section = global
	G.Global_key = cfg.Section("").Key("global_key").String()

	// Section = mysql
	G.Mysql_username = cfg.Section("mysql").Key("username").String()
	G.Mysql_password = cfg.Section("mysql").Key("password").String()
	G.Mysql_port = cfg.Section("mysql").Key("port").MustInt()
	G.Mysql_db = cfg.Section("mysql").Key("db").String()

	// Section = redis
	G.Redis_username = cfg.Section("redis").Key("username").String()
	G.Redis_password = cfg.Section("redis").Key("password").String()
	G.Redis_port = cfg.Section("redis").Key("port").MustInt()

}

/**
获得全局配置信息
 */
func GetConf()(*g){
 	return G
=======
	// Section = global
	global_key = cfg.Section("").Key("global_key").String()

	// Section = mysql
	mysql_username = cfg.Section("mysql").Key("username").String()
	mysql_password = cfg.Section("mysql").Key("password").String()
	mysql_port = cfg.Section("mysql").Key("port").MustInt()
	mysql_db = cfg.Section("mysql").Key("db").String()

	// Section = redis
	redis_username = cfg.Section("redis").Key("username").String()
	redis_password = cfg.Section("redis").Key("password").String()
	redis_port = cfg.Section("redis").Key("port").MustInt()
>>>>>>> 9f14ac8fb24acb5ef6109d89b707ec0be64a4e18
}