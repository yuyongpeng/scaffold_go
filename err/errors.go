package err

/**
错误信息的内容
*/
var StatusMsg = map[int]string{
	1001: "字符串类型错误",
	1002: "数据大于指定的值",
	1003: "请求没有响应",
}

func NewStatusError(errorNum int) error {
	return &StatusError{errorNum}
}

type StatusError struct {
	id int
}

func (e *StatusError) Error() string {
	return StatusMsg[e.id]
}























