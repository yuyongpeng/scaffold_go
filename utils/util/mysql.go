package util

import (
	"scaffold_go/config"
	"strconv"
)

/**
获得连接数据库的DSN
 */
func GetDsn()(dsn string){
	username := config.G.Mysql_username
	password := config.G.Mysql_password
	ip := config.G.Mysql_ip
	port := config.G.Mysql_port
	db := config.G.Mysql_db
	//dsn = "root:hard-chain2017@tcp(106.75.2.31:3306)/dht_msg?charset=utf8&parseTime=True&loc=Local"
	dsn = username + ":" + password + "@(" + ip + ":" + strconv.Itoa(port) + ")/" + db +"?charset=utf8&parseTime=True&loc=Local"
	return
}