package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"scaffold_go/conf"
	"scaffold_go/log"
	"strings"
	"sync"
	"time"
)

/**
用于连接数据库，关闭连接等
 */

type MysqlConnectPool struct {
	gdbx *gorm.DB

}
var once sync.Once
var instence *MysqlConnectPool
var gdb *gorm.DB
var err_db error
var cfg = conf.Cfg

var logger *logrus.Logger = log.Log

func GetInstence() *MysqlConnectPool {
	once.Do(func(){
		instence = &MysqlConnectPool{}
	})
	return instence
}

func (pool *MysqlConnectPool) InitDbPool() (success bool){
	dsn := strings.Join([]string{cfg.Mysql.Username,":",cfg.Mysql.Password,"@tcp(",cfg.Mysql.Ip,":",cfg.Mysql.Port, ")/",cfg.Mysql.Database,"?",cfg.Mysql.Param}, "")
	db, err_db := gorm.Open("mysql", dsn)
	if err_db != nil {
		logger.Error(err_db)
		return false
	}
	gdb = db
	pool.gdbx = db
	return true
}

func (pool *MysqlConnectPool) GetMysqlDB() (db_con *gorm.DB){
	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	gdb.DB().SetMaxIdleConns(cfg.Mysql.Maxidleconns)
	// SetMaxOpenConns 设置数据库的最大打开连接数。
	gdb.DB().SetMaxOpenConns(cfg.Mysql.Maxopenconns)
	// SetConnMaxLifetime 设置连接可以重用的最长时间。
	gdb.DB().SetConnMaxLifetime(time.Hour)
	//var k time.Duration = cfg.Mysql.Connmaxlifetime
	//gdb.DB().SetConnMaxLifetime(k)
	return gdb
}

func (m *MysqlConnectPool) Close(){
	gdb.Close()
}

func getDB() (db *gorm.DB){
	dsn := strings.Join([]string{cfg.Mysql.Username,":",cfg.Mysql.Password,"@tcp(",cfg.Mysql.Ip,":",cfg.Mysql.Port, ")/",cfg.Mysql.Database,"?",cfg.Mysql.Param}, "")
	db, err_db := gorm.Open("mysql", dsn)
	if err_db != nil {
		logger.Error(err_db)
	}
	// SetMaxIdleConns 设置空闲连接池中的最大连接数。
	db.DB().SetMaxIdleConns(cfg.Mysql.Maxidleconns)
	// SetMaxOpenConns 设置数据库的最大打开连接数。
	db.DB().SetMaxOpenConns(cfg.Mysql.Maxopenconns)
	// SetConnMaxLifetime 设置连接可以重用的最长时间。
	db.DB().SetConnMaxLifetime(time.Hour)
	return db
}











