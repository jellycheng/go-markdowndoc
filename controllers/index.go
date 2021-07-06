package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jellycheng/gosupport"
	"net/http"
	"time"
)

// 健康检查页 /health/index
func Health(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"code":0, "msg":"success", "data":new(struct{})})
}

// 健康检查页 /health/env2data
func HealthEnvData(c *gin.Context)  {
	retData := make(map[string]interface{})
	envFile := globalCfg.GetString("ENV_FILE");
	retData["env_file"] = envFile

	for k,v := range globalEnv.Data {
		retData[k] = v
	}
	c.JSON(http.StatusOK, gin.H{"code":0, "msg":"success", "data":retData})
}

// 服务启动信息
func StartInfo(c *gin.Context)  {
	var ret = make(map[string]interface{})
	ret["cur_time"] = gosupport.TimeNow2Format("2006-01-02 15:04:05")
	if start_time,ok:=globalCfg.Get("SERVER_START_TIME");ok {
		if v,ok:=start_time.(time.Time);ok {
			ret["server_start_time"] = v.Format("2006-01-02 15:04:05")
			ret["server_run_time"] = "服务运行了 " + gosupport.AlreadyTimeStr(v)
		} else {
			ret["server_start_time"] = start_time
		}
	}

	c.JSON(http.StatusOK, gin.H{"code":0, "msg":"success", "data":ret})

}

// 首页
func Index(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{"code":0, "msg":"success", "data":new(struct{})})
}


