package main

import (
	"github.com/go-ini/ini"
	"fmt"
	"os"
)
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
	cfg,err := ini.Load(confFile)
	if err != nil{
		fmt.Println("Fail to read file :", confFile)
		os.Exit(1)
	}
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
}