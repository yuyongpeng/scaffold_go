package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
	"log"
	"fmt"
	"time"
<<<<<<< HEAD
	)
=======
)
>>>>>>> 9f14ac8fb24acb5ef6109d89b707ec0be64a4e18

type MysqlConnectPool struct {

}
var once sync.Once
var instence *MysqlConnectPool
var gdb *gorm.DB
var err_db error

func GetInstence() *MysqlConnectPool {
	once.Do(func(){
		instence = &MysqlConnectPool{}
	})
	return instence
}

func (m *MysqlConnectPool) InitDbPool() (sucess bool){
	db, err_db := gorm.Open("mysql", "root:hard-chain2017@tcp(106.75.2.31:3306)/dht_msg?charset=utf8&parseTime=True&loc=Local")
	fmt.Println(err_db)
	if err_db != nil {
		log.Fatal(err_db)
		return false
	}
	gdb = db
	//db.Close()
	return true
}

func (m *MysqlConnectPool) GetMysqlDB() (db_con *gorm.DB){
	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	gdb.DB().SetMaxIdleConns(10)
	// SetMaxOpenConns 设置数据库的最大打开连接数。
	gdb.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置连接可以重用的最长时间。
	gdb.DB().SetConnMaxLifetime(time.Hour)
	return gdb
}

func (m *MysqlConnectPool) Close(){
	gdb.Close()
}













