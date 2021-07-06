package routers

import (
	"github.com/gin-gonic/gin"
	"markdowndoc/controllers"
)

func RegisterRouters(r *gin.Engine) {
	//健康检查页,支持任意请求方式
	r.Any("/health/index", controllers.Health)
	//首页
	r.GET("/", controllers.Index)
	r.Any("/favicon.ico", func(c *gin.Context) {
		c.String(200, "")
	})
	//服务启动信息
	r.GET("/start/info", controllers.StartInfo)

	//提供C端api接口
	RegisterApiRouters(r)

	//404
	r.NoRoute(controllers.NoRoute)


}
