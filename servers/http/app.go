package http

import (
	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/servers/http/users"
)

func Start() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	users.RouteRegister(v1.Group("/auth"))
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
