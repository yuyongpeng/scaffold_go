package database

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	"time"
)

/**
存放数据库表的模型
*/

type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
	Num      int    `gorm:"AUTO_INCREMENT"` // 自增

	CreditCard CreditCard // One-To-One (拥有一个 - CreditCard表的UserID作外键)
	Emails     []Email    // One-To-Many (拥有多个 - Email表的UserID作外键)

	BillingAddress   Address // One-To-One (属于 - 本表的BillingAddressID作外键)
	BillingAddressID sql.NullInt64

	ShippingAddress   Address // One-To-One (属于 - 本表的ShippingAddressID作外键)
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`                         // 忽略这个字段
	Languages []Language `gorm:"many2many:user_languages;"` // Many-To-Many , 'user_languages'是连接表
}

type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // 外键 (属于), tag `index`是为该列创建索引
	Email      string `gorm:"type:varchar(100);unique_index"` // `type`设置sql类型, `unique_index` 为该列设置唯一索引
	Subscribed bool
}

type Address struct {
	ID       int
	Address1 string         `gorm:"not null;unique"` // 设置字段为非空并唯一
	Address2 string         `gorm:"type:varchar(100);unique"`
	Post     sql.NullString `gorm:"not null"`
}

type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // 创建索引并命名，如果找到其他相同名称的索引则创建组合索引
	Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

////////////////////////// [table: creditcard] //////////////////////////
/**
CREATE TABLE `creditcard` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userid` int(11) DEFAULT NULL,
  `number` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
*/
type Model struct {
	ID        uint       `gorm:"primary_key"`
	CreatedAt time.Time  `gorm:"type:datetime;Column:createat"`
	UpdatedAt time.Time  `gorm:"type:datetime;Column:updateat"`
	DeletedAt *time.Time `sql:"index"`
}
type CreditCard struct {
	//gorm.Model
	Id     int    `gorm:"type:int;AUTO_INCREMENT;NOT NULL;UNIQUE"`
	UserID uint   `gorm:"type:int;Column:userid"`
	Number string `gorm:"type:varchar(100);Column:number"`
}

func (CreditCard) TableName() string {
	return "creditcard"
}

////////////////////////// [table: cp_soldier] //////////////////////////
type Soldier struct {
	Id               int       `gorm:"type:int(10);column:id;AUTO_INCREMENT;NOT NULL;UNIQUE"`
	Encrypted_string string    `gorm:"type:varchar(700);Column:encrypted_string;NOT NULL;UNIQUE"`
	Verification_num int       `gorm:"type:int(11);Column:verification_num;default:0"`
	Update_time      time.Time `gorm:"type:datetime;Column:update_time"`
	Create_time      time.Time `gorm:"type:datetime;Column:create_time"`
}

func (Soldier) TableName() string {
	return "cp_soldiers"
}
