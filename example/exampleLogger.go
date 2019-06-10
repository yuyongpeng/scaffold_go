/**
Author:       yuyongpeng@hotmail.com
Github:       https://github.com/yuyongpeng/
Date:         2019-06-10 13:56:56
LastEditors:
LastEditTime: 2019-06-10 13:56:56
Description:
*/
package main

import "github.com/sirupsen/logrus"
import "scaffold_go/log"

var logger *logrus.Logger = log.Log

func main() {
logger.WithFields(logrus.Fields{
	"common": "this is a common field",
	"other": "I also should be logged always",
}).Info("xxxxxxxxx")

}
