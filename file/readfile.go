/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-11 20:50:05
LastEditors:
LastEditTime: 2019-06-11 20:50:05
Description:
*/
package file

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"scaffold_go/crypto"
	"scaffold_go/database"
	"scaffold_go/errors"
	"strings"
)

/**
解析文件导入数据库中
 */
func ReadFile(filePath string, handle func(string) (string, error)) (errorLines []string, success int, retEr error) {
	var lines []string
	success = 0
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return lines, success, err
	}
	bufReader := bufio.NewReader(f)
	for {
		line, err := bufReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return lines, success, nil
			}
			return lines, success, err
		}
		line = strings.TrimSpace(line)
		// 跳过空行
		if line == "" {
			continue
		}
		// 处理一行数据
		if line, e2 := handle(line); e2 != nil {
			lines = append(lines, line)
		}else{
			success += 1
		}
	}
	return lines, success, nil
}

/*

解析csv文件
CSV 文件的格式："姓名","身份证号","出生年月","兵种","专业","退役号"
示例："光头强","362101198721010087","1982-01-08","步兵","列兵","110982091"

nodejs中对应的字段说明：
name : "光头强",
id_number : '362101198721010087',
birth_date : "1982-10-17",
major : "步兵",
arms : "列兵",
demobilized_number : "110982091",

如果解析成功，返回的是加密后的字符串
如果解析失败，返回的是 json 行

*/
func PaseCsvToMysql(line string) (encryptedString string, e error) {
	ioreader := strings.NewReader(line)
	reader := csv.NewReader(ioreader)
	if recorde, er := reader.Read(); er == nil {
		name := recorde[0]
		id_number := recorde[1]
		birth_date := recorde[2]
		major := recorde[3]
		arms := recorde[4]
		demobilized_number := recorde[5]
		encryptedString = strings.Join([]string{name, id_number, birth_date, major, arms, demobilized_number}, "")
		hash := crypto.GetKeccakHash(encryptedString)
		crud := &database.Crud{}
		if insertErr := crud.InsertEncryptedString(hash); insertErr != nil{
			return line, &errors.StatusError{Id: 1005}
		}else{
			return encryptedString, nil
		}
	} else {
		return line, &errors.StatusError{Id: 1004}
	}
}
