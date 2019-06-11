/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-11 20:07:23
LastEditors:
LastEditTime: 2019-06-11 20:07:23
Description:
*/
package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/ebfe/keccak"
	"golang.org/x/crypto/sha3"
)

func main(){
	k := keccak.New256()
	k.Reset()
	str := []byte("yuyongpeng")
	k.Write(str)
	a := k.Sum(nil)
	//fmt.Print(string(a))
	fmt.Println(hex.EncodeToString(a))

	sa := sha512.New()
	kk := sa.Sum(str)
	fmt.Println(hex.EncodeToString(kk))

	a3 := sha3.New384()
	k2 := a3.Sum(str)
	fmt.Println(hex.EncodeToString(k2))

}