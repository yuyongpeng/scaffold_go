package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
/**
存放对数据库的操作，增加、删除、查询等
 */
type Crud struct {

}

func (crud *Crud) getUsers(gdb *gorm.DB){
	gdb.AutoMigrate(&Email{})
}

func (crud *Crud) GetCreditCard(){
	db := getDB()
	db.LogMode(true)
	defer db.Close()
	//db.Exec("insert into creditcard (userid, number) values(123,'sdfsdfdfs')")
	var cred CreditCard
	//db.First(&cred, 2)
	db.Where("id = ?", 2).Find(&cred)
	//db.Model(&cred).Update("Number", "111111")
	fmt.Println(cred.Number)
	//db.Close()
}


