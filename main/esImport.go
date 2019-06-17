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
	"flag"
	"scaffold_go/database"
	"scaffold_go/elastic"
	"strings"
)

func main() {
	var types string
	flag.StringVar(&types, "t", "person", "导入数据的类型（job和person）")
	flag.Parse()
	types = strings.TrimSpace(types)
	crud := &database.Escrud{}
	if types == "job" {
		cu := crud.GetJobsCount()
		cu = 20
		var inc int = 10
		for i := 0; i <= cu; i = i + inc {
			jobs := crud.GetJobs(i, i+inc)
			elastic.ImportJobs(jobs)
		}
	}
	if types == "person" {
		cu := crud.GetPersonsCount()
		cu = 20
		var inc int = 10
		for i := 0; i <= cu; i = i + inc {
			person := crud.GetPersons(i, i+inc)
			elastic.ImportPersons(person)
		}
	}
}
