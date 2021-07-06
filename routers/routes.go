package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jellycheng/gosupport"
	"markdowndoc/controllers"
)

var globalenv *gosupport.DataManage

func init() {
	globalenv = gosupport.NewGlobalEnvSingleton()
}

func RegisterRouters(r *gin.Engine) {
	//健康检查页,支持任意请求方式
	r.Any("/health/index", controllers.Health)
	//首页
	r.GET("/", controllers.HtmlIndex)
	r.Any("/favicon.ico", func(c *gin.Context) {
		c.String(200, "")
	})
	//服务启动信息
	r.GET("/start/info", controllers.StartInfo)
	r.LoadHTMLGlob("./views/*")

	//提供C端api接口
	RegisterApiRouters(r)

	// 静态资源
	staticDir := globalenv.GetString("STATIC_DIR")
	if staticDir == "" {
		staticDir = "./static/"
	}
	r.Static("/static", staticDir)

	//404
	r.NoRoute(controllers.NoRoute)


}
