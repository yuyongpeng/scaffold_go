/**
      ___           ___           ___
     /\__\         /\__\         /\  \
    /:/ _/_       /:/  /        /::\  \
   /:/ /\  \     /:/  /        /:/\:\  \
  /:/ /::\  \   /:/  /  ___   /:/ /::\  \
 /:/_/:/\:\__\ /:/__/  /\__\ /:/_/:/\:\__\
 \:\/:/ /:/  / \:\  \ /:/  / \:\/:/  \/__/
  \::/ /:/  /   \:\  /:/  /   \::/__/
   \/_/:/  /     \:\/:/  /     \:\  \
     /:/  /       \::/  /       \:\__\
     \/__/         \/__/         \/__/
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-24 07:34:22
LastEditors:
LastEditTime: 2019-06-24 07:34:22
Description:
*/
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Student struct {
	Name string `学生姓名`
	Age  int    `a:"1111"b:"2222"`
}

func main() {
	s := Student{}
	rt := reflect.TypeOf(s)
	fieldName, ok := rt.FieldByName("Name")
	if ok {
		fmt.Println(fieldName)
	}
	fieldAge, ok2 := rt.FieldByName("Age")
	if ok2 {
		fmt.Println(fieldAge.Tag.Get("a")) // 获得结构体 tag的属性
		fmt.Println(fieldAge.Tag.Get("b"))
	}
	fmt.Println("type Name: " + rt.Name())                       // Student
	fmt.Println("type NumField: " + strconv.Itoa(rt.NumField())) // 2
	fmt.Println("type PkgPath: " + rt.PkgPath())                 // main
	fmt.Println("type String: " + rt.String())                   // main.Student

	fmt.Println(rt.Kind().String()) // struct

	for i := 0; i < rt.NumField(); i++ { // 循环取出结构体的属性
		fmt.Println(rt.Field(i).Name) // Name  Age
	}

	sc := make([]int, 10)
	sc = append(sc, []int{1, 2, 3}...)
	sct := reflect.TypeOf(sc)
	scet := sct.Elem()  //////////

	fmt.Println("slice element type.Name() = " + scet.Name())		// int
	fmt.Println("slice element type.NumMethod() = ", scet.NumMethod())		// 0
	fmt.Println("slice element scet type.PkgPath() = ", scet.PkgPath())	// 没有输出
	fmt.Println("slice element sct type.PkgPath() = ", sct.PkgPath())		// 没有输出

	fmt.Println("slice element type.Kind() = ", scet.Kind())		// int
	fmt.Printf("slice element type.Kind() = %d\n", scet.Kind())	// 2
	fmt.Println("slice element type.String() = ", scet.String())	// int


	key := "0d1ea7347d7a438ebd7d8a9a6da8f678"
	secrt := "29FE2DA222B5579F4F1B1EE199E8F547"
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)
	timestampStr := strconv.FormatInt(timestamp, 10)
	data := strings.Join([]string{key, timestampStr, secrt}, "")
	// b2aa284aa7c2bbf35e5967aae17b9e8b
	//fmt.Println(hex.EncodeToString())  c950b040bcfeef8bdca8359d1286a5e7
	m5 := md5V(data)
	fmt.Println(strings.ToUpper(m5))

}

func md5V(str string) string  {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}