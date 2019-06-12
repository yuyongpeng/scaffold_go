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
Date:         2019-06-12 12:56:38
LastEditors:
LastEditTime: 2019-06-12 12:56:38
Description: 处理iris的数据
*/
package service

import "scaffold_go/file"

/**
将 CSV 文件解析后导入到数据库中
 */
func CsvService(csvFile string) (errorLine []string, success int) {
	errorLine, success, _ = file.ReadFile(csvFile, file.PaseCsvToMysql)
	return errorLine, success
}
