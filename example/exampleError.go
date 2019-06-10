/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/scaffold_go
Date:         2019-06-09 10:38:18
LastEditors:
LastEditTime: 2019-06-09 10:38:18
Description:
*/
package main

import (
	"fmt"
	"math"
	"scaffold_go/err"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, err.NewStatusError(1002)
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print("area of circle %0.2f", area)
}
