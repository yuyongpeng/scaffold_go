package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"scaffold_go/errors"
	"scaffold_go/log"
	"time"
)

/**
存放对数据库的操作，增加、删除、查询等
*/
type Crud struct {
}

func (crud *Crud) getUsers(gdb *gorm.DB) {
	gdb.AutoMigrate(&Email{})
}

func (crud *Crud) GetCreditCard() {
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

/**
将生成的hash串插入数据库中
*/
func (crud *Crud) InsertEncryptedString(encryptedString string) error {
	var logger *logrus.Logger = log.Log
	logger.Info(encryptedString)
	db := getDB()
	//db.LogMode(true)
	defer db.Close()
	var soldier Soldier
	//fmt.Printf("hash : %s", encryptedString)
	tx := db.Where(Soldier{Encrypted_string: encryptedString}).
		Attrs(Soldier{Verification_num: 0, Update_time: time.Now(), Create_time: time.Now()}).
		FirstOrCreate(&soldier)
	if err := tx.Error; err != nil {
		logger.Errorln(err)
		return &errors.StatusError{Id: 1001}
	} else {
		return nil
	}
}


