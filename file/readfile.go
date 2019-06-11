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
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadFile(filePath string, handle func(string)) error {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return err
	}
	bufReader := bufio.NewReader(f)
	for {
		line, err := bufReader.ReadString('\n')
		line = strings.TrimSpace(line)
		handle(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}

// 解析csv文件
func PaseCsv(line string) {
	ioreader := strings.NewReader(line)
	reader := csv.NewReader(ioreader)
	if recorde, err := reader.Read(); err == nil {
		fmt.Println(recorde[1])
	}
}
