package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jellycheng/gosupport"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

//日志中间件
func LogMiddle(c *gin.Context) {
	logrus.Debug("执行LogMiddle中间件")
	//设置日志输出多端 0644、0755
	if globalCfg.GetString("SERVER_LOGRUS_LASTFILE_DATE")!=gosupport.TimeNow2Format("20060102") {
		targetDir := fmt.Sprintf("%s%s",globalEnv.GetString("LOG_DIR"),globalEnv.GetString("APP_NAME"))
		if !gosupport.IsDir(targetDir) {
			os.MkdirAll(targetDir, os.ModePerm)
		}
		logFileName := fmt.Sprintf("%s/%s.%s.log", targetDir,globalEnv.GetString("APP_NAME"),gosupport.TimeNow2Format("2006-01-02"))
		writerF, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err == nil {
			logrus.SetOutput(io.MultiWriter(os.Stdout, writerF))
		} else {
			logrus.Errorf("create file %s failed: %v", logFileName, err)
			logrus.SetOutput(io.MultiWriter(os.Stdout))
		}

		globalCfg.Set("SERVER_LOGRUS_LASTFILE_DATE", gosupport.TimeNow2Format("20060102"))
	}

}

