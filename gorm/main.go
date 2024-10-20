package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"scaffold_go/gorm/dal/query"
)

const MySQLDSN = "root:123456@(172.22.135.106:3306)/testgen?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open(MySQLDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	query.SetDefault(db)
	//查询User年龄为18的
	userM := query.User
	user, _ := userM.Where(userM.Age.Eq(18)).Preload(userM.Address).Find()
	fmt.Print(user[0].Address)

	u := query.User
	a := query.Address
	ux, _ := u.Select(u.ALL, a.ALL).LeftJoin(a, a.UID.EqCol(u.ID)).Find()
	fmt.Println(ux)

}
