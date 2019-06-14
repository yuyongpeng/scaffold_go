/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-10 14:33:47
LastEditors:
LastEditTime: 2019-06-10 14:33:47
Description:
*/
package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"scaffold_go/database"
)

func main() {
	//cfg := conf.Cfg
	//fmt.Printf("name : %s\n", cfg.Mysql.Username)
	//// root:1q2w3e4r@tcp(127.0.0.1:3306)/scaffold?charset=utf8mb4&parseTime=True&loc=Local
	//dsn := strings.Join([]string{cfg.Mysql.Username,":",cfg.Mysql.Password,"@tcp(",cfg.Mysql.Ip,":",cfg.Mysql.Port, ")/",cfg.Mysql.Database,"?",cfg.Mysql.Param}, "")
	//fmt.Println(dsn)
	//db, err := gorm.Open("mysql", dsn)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//defer db.Close()
	//db.Exec("insert into creditcard (userid, number) values(123,'sdfsdfdfs')")

	crud := &database.Escrud{}
	//for {
	//	//crud.GetCreditCard()
	//	//time.Sleep(1 * time.Second)
	//}

	cu := crud.GetJobsCount()
	fmt.Println(cu)

	crud.GetJobs(1,100)

}
