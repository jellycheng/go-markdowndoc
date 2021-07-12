package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/curl"
	"github.com/russross/blackfriday/v2"
	"net/http"
	"regexp"
	"strings"
	"time"
	"html/template"
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

func HtmlIndex(c *gin.Context)  {
	mdDir := curl.TrimPath(globalEnv.GetString("MD_DOC_ROOT_PATH"),2) + "/"
	projectName := "demo"
	pname := c.Query("pname")
	if pname != "" {
		if isMatch,err := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9_-]*$", pname);err == nil && isMatch {
			projectName = pname
		} else {
			c.String(200, "文档不存在")
			return
		}

	}

	leftContent := ""
	leftFile := fmt.Sprintf("%s/%s/SUMMARY.md", mdDir, projectName)
	if gosupport.IsFile(leftFile) {
		leftFileCon, _ := gosupport.FileGetContents(leftFile)
		leftFileCon = strings.ReplaceAll(leftFileCon, "(docs/", "(?pname="+projectName+"&md=")
		leftContentByte := blackfriday.Run([]byte(leftFileCon))
		leftContent = string(leftContentByte)
	}

	md := c.Query("md")
	if md == "" {
		md = "index"
	}
	rightContent := "文档不存在"
	rightFile := fmt.Sprintf("%s/%s/docs/%s.md", mdDir, projectName, strings.TrimRight(md, ".md"))
	if gosupport.IsFile(rightFile) {
		rightFileCon,_ := gosupport.FileGetContents(rightFile)
		rightContentByte := blackfriday.Run([]byte(rightFileCon))
		rightContent = string(rightContentByte)
	}

	c.HTML(200, "layout_default.html", struct {
		LeftContent template.HTML
		RightContent template.HTML
	}{
		LeftContent: template.HTML(leftContent),
		RightContent: template.HTML(rightContent),
	})

}
