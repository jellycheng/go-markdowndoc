package cmd

import (
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/env"
	"github.com/jellycheng/gosupport/xversion"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"fmt"
	mylog "markdowndoc/library/log"
	"time"
)

var (
	START_TIME time.Time   //启动时间
	globalCfg *gosupport.DataManage
	globalEnv *gosupport.DataManage
)

func init()  {
	START_TIME = time.Now()
	xversion.Version = "v1.0.0"
	globalCfg = gosupport.NewGlobalCfgSingleton()
	globalEnv = gosupport.NewGlobalEnvSingleton()
	globalCfg.Set("SERVER_START_TIME", START_TIME) //服务启动时间
	globalCfg.Set("APP_CODE_VERSION", xversion.Version) //代码版本

}

//cmd公共的方法
func commonActionInit(envFile string)  {
	//解析全局配置文件
	err := env.LoadEnv2DataManage(envFile)
	if err!=nil {
		fmt.Println(err.Error())
	}
	globalEnv = gosupport.NewGlobalEnvSingleton()
	globalCfg.Set("app_name", globalEnv.GetString("APP_NAME"))

	//设置日志格式
	logrus.SetFormatter(new(mylog.LogFormatter))
	//设置日志级别
	logLevel, err := logrus.ParseLevel(globalEnv.GetString("LOG_LEVEL"))
	if err != nil {
		fmt.Println(err.Error())
	}
	logrus.SetLevel(logLevel)
	//设置日志输出多端
	targetDir := fmt.Sprintf("%s%s",globalEnv.GetString("LOG_DIR"),globalEnv.GetString("APP_NAME"))
	if !gosupport.IsDir(targetDir) {
		os.MkdirAll(targetDir, os.ModePerm)
	}
	logFileName := fmt.Sprintf("%s/%s.%s.log", targetDir,globalEnv.GetString("APP_NAME"),gosupport.TimeNow2Format("2006-01-02"))
	writerF, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Printf("create file %s failed: %v", logFileName, err)
		os.Exit(0)
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, writerF))
	globalCfg.Set("SERVER_LOGRUS_LASTFILE_DATE", gosupport.TimeNow2Format("20060102"))


}
