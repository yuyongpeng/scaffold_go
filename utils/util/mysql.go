package util

import (
	"fmt"
	"scaffold_go/conf"
	"reflect"
	"strconv"
)

/**
获得连接数据库的DSN
 */
func GetDsn()(dsn string){
	username := conf.G.Mysql_username
	password := conf.G.Mysql_password
	ip := conf.G.Mysql_ip
	port := conf.G.Mysql_port
	db := conf.G.Mysql_db
	fmt.Println(reflect.TypeOf(port))
	//dsn = "root:hard-chain2017@tcp(106.75.2.31:3306)/dht_msg?charset=utf8&parseTime=True&loc=Local"
	dsn = username + ":" + password + "@(" + ip + ":" + strconv.Itoa(port) + ")/" + db +"?charset=utf8&parseTime=True&loc=Local"
	fmt.Println(dsn)
	return
}