/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-09 10:58:16
LastEditors:
LastEditTime: 2019-06-09 10:58:16
Description:
*/
package test

import (
	"fmt"
	"github.com/spf13/viper"
	"scaffold_go/conf"
)

var v = viper.New()

func GetViperPro(){
	e := v.GetString("eyes")
	fmt.Printf("aaaa: %s", e)
}

func Test(){
	vp := conf.Vp
	name := vp.Get("name")
	fmt.Printf("xxxxx: %s", name)
}

func main(){
	var G = conf.G
	fmt.Println(G.Formater)
}
