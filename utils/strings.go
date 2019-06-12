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
Date:         2019-06-12 15:02:38
LastEditors:
LastEditTime: 2019-06-12 15:02:38
Description:
*/
package utils

import (
	"math/rand"
	"time"
)

const (
	KC_RAND_KIND_NUM   = 0  // 纯数字
	KC_RAND_KIND_LOWER = 1  // 小写字母
	KC_RAND_KIND_UPPER = 2  // 大写字母
	KC_RAND_KIND_ALL   = 3  // 数字、大小写字母
)

/**
  随机字符串
  fmt.Println("num:   " + string(Krand(16, KC_RAND_KIND_NUM)))
  fmt.Println("lower: " + string(Krand(16, KC_RAND_KIND_LOWER)))
  fmt.Println("upper: " + string(Krand(16, KC_RAND_KIND_UPPER)))
  fmt.Println("all:   " + string(Krand(16, KC_RAND_KIND_ALL)))
 */
func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i :=0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base+rand.Intn(scope))
	}
	return result
}