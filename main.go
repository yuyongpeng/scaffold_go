package main
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"scaffold_go/database"
	"scaffold_go/conf"
	"fmt"
	"scaffold_go/utils/util"
)

type Product struct {
	gorm.Model
	Code string
	Price int
}
func main() {
	conf.Initial("./conf.ini")
	G := conf.GetConf()
	fmt.Println(G.Mysql_username)
	dsn := util.GetDsn()
	fmt.Println("xx  ",conf.G.Mysql_username)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code:"124", Price: 12})

	var product Product
	db.First(&product, 1)
	db.First(&product, "code=?", "124")

	db.Model(&product).Update("Price", 2000)
	db.Delete(&product)

	instance := database.GetInstence()
	conPool := instance.InitDbPool()
	if conPool == false{
		panic("connect error")
	}
	gdb := instance.GetMysqlDB()
	gdb.AutoMigrate(&database.Email{})

	}
