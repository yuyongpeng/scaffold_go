/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/scaffold_go
Date:         2019-06-09 10:38:18
LastEditors:
LastEditTime: 2019-06-09 10:38:18
Description:
*/
package main

import (
	"fmt"
	"scaffold_go/conf"
)

func main() {
	conf.NewVp()
	fmt.Printf("ddd:%s", conf.Cfg.Name)
	cfg := conf.Cfg
	fmt.Printf("name : %s", cfg.Mysql.Username)
}
