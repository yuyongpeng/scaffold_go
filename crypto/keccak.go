/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-11 20:37:43
LastEditors:
LastEditTime: 2019-06-11 20:37:43
Description:
*/
package crypto

import (
	"encoding/hex"
	"github.com/ebfe/keccak"
)

// 生成字符串的keccackHash  256
func GetKeccakHash(cop string) ( hash string){
	keak := keccak.New256()
	keak.Reset()
	keak.Write([]byte(cop))
	hashByte := keak.Sum(nil)
	return hex.EncodeToString(hashByte)
}


