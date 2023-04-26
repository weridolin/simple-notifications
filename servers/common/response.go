package common

import "github.com/gin-gonic/gin"

func HttpResponse(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	return
}
