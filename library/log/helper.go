package log

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"runtime"
)

// 带有trace的日志打印
// WithContext(reqContext).Info("foo", bar, ....)
// WithContext(reqContext).Infof("foo:%v", bar, ....)
// WithContext(reqContext).Error("error message", bar, ....)
// WithContext(reqContext).Errorf("error message:%v", bar, ....)
func WithContext(c *gin.Context) *logrus.Entry {
	pc, _, _, _ := runtime.Caller(2)
	f := runtime.FuncForPC(pc)
	if c == nil {
		return logrus.WithFields(map[string]interface{}{
			"user_id":       "",
			"enterprise_id": "",
			"traceid":       "",
			"func":          f.Name(),
		})
	}
	return logrus.WithFields(map[string]interface{}{
		"user_id":       c.GetInt64("login_userid"),
		"enterprise_id": c.GetString("enterprise_id"),
		"traceid":       c.GetString("traceid"),
		"func":          f.Name(),
	})
}

