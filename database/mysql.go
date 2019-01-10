package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"sync"
	"scaffold_go/log"
	"scaffold_go/utils/util"
	"time"
)
type MysqlConnectPool struct {

}
var once sync.Once
var instence *MysqlConnectPool
var gdb *gorm.DB
var err_db error

//var logger = log.New()
var logger *logrus.Logger = log.New()

func GetInstence() *MysqlConnectPool {
	once.Do(func(){
		instence = &MysqlConnectPool{}
	})
	return instence
}

func (m *MysqlConnectPool) InitDbPool() (sucess bool){
	//var logger = log.New()
	logger.Info("mysql database")
	dsn := util.GetDsn()
	db, err_db := gorm.Open("mysql", dsn)
	if err_db != nil {
		//log.Fatal(err_db)
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













