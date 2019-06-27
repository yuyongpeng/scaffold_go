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
Date:         2019-06-24 10:41:34
LastEditors:
LastEditTime: 2019-06-24 10:41:34
Description:
*/
package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

/*
生成md5的加密串
*/
func GenMd5V1(original string) (md5Str string) {
	data := []byte(original)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

/*
生成md5的加密串
*/
func GenMd5V2(original string) (md5Str string) {
	h := md5.New()
	h.Write([]byte(original))
	return hex.EncodeToString(h.Sum(nil))
}
