/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-24 10:48:37
LastEditors:
LastEditTime: 2019-06-24 10:48:37
Description:
*/
package crypto

import (
	"testing"
)

func TestGenMd5V1(t *testing.T){
	orig := "1q2w3e4r5t6y7u8i9o0p"
	md5Str := GenMd5V1(orig)
	//fmt.Println(md5Str)
	if md5Str != "c6b419d72d1664762ad3d5c500566c69"{
		t.Errorf("GenMd5v1 方法生成说句错误 Error")
	}
}