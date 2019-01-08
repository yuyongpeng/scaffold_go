package database
<<<<<<< HEAD

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	)

type Crud struct {

}

func (crud *Crud) getUsers(gdb *gorm.DB){
	gdb.AutoMigrate(&Email{})
}
=======
>>>>>>> 9f14ac8fb24acb5ef6109d89b707ec0be64a4e18
