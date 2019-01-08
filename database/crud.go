package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	)

type Crud struct {

}

func (crud *Crud) getUsers(gdb *gorm.DB){
	gdb.AutoMigrate(&Email{})
}
