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
Date:         2019-06-14 11:50:57
LastEditors:
LastEditTime: 2019-06-14 11:50:57
Description:
*/
package main

import (
	"fmt"
	"scaffold_go/database"
)

func main() {
	crud := &database.Escrud{}
	cu := crud.GetJobsCount()
	var inc int = 2
	for i:=0; i<=cu; i = i+inc{
		jobs := crud.GetJobs(i,i + inc)
		for _, value := range jobs{
			fmt.Println(value)
		}
	}
}
