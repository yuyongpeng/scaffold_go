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
Date:         2019-06-13 16:04:23
LastEditors:
LastEditTime: 2019-06-13 16:04:23
Description:
*/
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	const jsonStream = `
	{"Name": "Ed", "Text": ["a","b"]}
`
	var r map[string]interface{}
	dec := json.NewDecoder(strings.NewReader(jsonStream))
	dec.Decode(&r)
	for key, value := range r{
		fmt.Println("name : " + key)
		fmt.Println(reflect.TypeOf(value))
		x := value.([]string)
		fmt.Println(x)
	}

	var obj interface{} = []string{"a", "b", "c"}
	arr := obj.([]string)
	fmt.Printf(arr[0])

}
