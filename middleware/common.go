package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	myutils "markdowndoc/library/utils"
	"strconv"
)

//通用逻辑处理，必须调用
func CommonMiddle(c *gin.Context) {
	logrus.Debug("执行CommonMiddle中间件")
	//获取TRACE-ID，没有生成一个TRACE-ID
	headerTraceid := c.Request.Header.Get("TRACE-ID")
	if headerTraceid != "" && headerTraceid != "0" {
		c.Set("traceid", headerTraceid)
	} else {
		headerTraceid = c.Request.Header.Get("X-TRACE-ID")
		if headerTraceid != "" && headerTraceid != "0" {
			c.Set("traceid", headerTraceid)
		} else {
			c.Set("traceid", myutils.GenerateTraceId())
		}
	}

	//登录者用户ID
	userid := c.Request.Header.Get("X-USERID")

	if intUserId64, err := strconv.ParseInt(userid, 10, 64); err == nil {
		c.Set("login_userid", intUserId64)
	} else {
		c.Set("login_userid", 0)
	}

}

