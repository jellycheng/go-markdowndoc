package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)


func NoRoute(c *gin.Context) {
	curUrl := c.Request.URL.Path
	logrus.Error(fmt.Sprintf("请求地址不存在:%s", curUrl))

	c.JSON(200, map[string]interface{}{"code":404,
											"msg":"地址不存在",
											"data":map[string]interface{}{"url":curUrl},
		 })

}

