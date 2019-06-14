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
Description:  错误处理

			调用方式:
			var e error
			e = errors.StatusError{Id: 1001}

*/
package errors

/**
错误信息的内容
*/
var StatusMsg = map[int]string{
	1001: "字符串类型错误",
	1002: "数据大于指定的值",
	1003: "请求没有响应",
	1004: "CSV解析错误",
	1005: "数据插入数据库失败",
	1006: "CSV的字段长度不对",
	1007: "日期格式不对",
	1008: "职位信息插入elasticsearch失败",
}

func NewStatusError(errorNum int) error {
	return &StatusError{errorNum}
}

type StatusError struct {
	Id int
}

func (e *StatusError) Error() string {
	return StatusMsg[e.Id]
}























