package config

//错误代号配置，常量首字母大写，驼峰式
const (
	Success             = 0
	Fail                = 1
	ErrorRecordNotExist = 200
	ErrorParam          = 201
	ErrorIllegalRequest = 202
	ErrorNotLogin       = 10000
	ErrorDefault        = 10001 + iota
	ErrorDBNotFound
	ErrorRequestParam
	ErrorMysqlConnection
	ErrorRedisConnection
	ErrorAddFailed
	ErrorActionFailed
)

var codeMsg = map[int]string{
	Success:                      "success",
	Fail:                         "fail",
	ErrorNotLogin:                "未登录",
	ErrorRecordNotExist:          "记录不存在",
	ErrorParam:                   "参数不合法",
	ErrorIllegalRequest:          "非法请求",
	ErrorDefault:                 "系统错误",
	ErrorDBNotFound:              "数据库不存在",
	ErrorRequestParam:            "请求参数错误",
	ErrorMysqlConnection:         "mysql连接错误",
	ErrorRedisConnection:         "redis连接错误",
	ErrorAddFailed:               "新增失败",
	ErrorActionFailed:            "操作失败",
}

//通过代号获取错误信息
func GetMsg4Code(code int) string {
	if m, ok := codeMsg[code]; ok == true {
		return m
	}
	return codeMsg[ErrorDefault]
}

