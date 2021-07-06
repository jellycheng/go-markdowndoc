package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/uuid"
	"runtime"
	"time"
)

// 创建gin.Context，一般用于job、kafka等
func NewGinCtx() *gin.Context {
	c := new(gin.Context)
	c.Set("traceid", GenerateTraceId())
	c.Set("branch_name", "feature_mddoc_20200827") // 开发环境会有用
	return c
}


// 获取正在运行的函数信息
func GetFuncInfo(skip int) (funcName string, fileName string, line int, ok bool) {
	pc, file, line, ok := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	return f.Name(), file, line, ok
}

//生成新的traceID
func GenerateTraceId() string {
	tid := uuid.GenerateUUID(time.Now())
	return gosupport.Md5V1(tid)
}

