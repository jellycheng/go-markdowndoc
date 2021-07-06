package routers

import (
	"github.com/gin-gonic/gin"
	"markdowndoc/controllers"
)

//提供C端api接口，移动api接口
func RegisterApiRouters(r *gin.Engine) {

	api := r.Group("/api/v1")
	{
		api.GET("/", controllers.Index)
	}

}

