package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/weridolin/simple-vedio-notifications/servers/http/middlewares"
	"github.com/weridolin/simple-vedio-notifications/servers/http/users"
)

func Start() {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "v1 api",
		})
	})

	users.UnAuthorizationRouteRegister(v1.Group("/auth"))

	v1.Use(middlewares.AuthorizationMiddleware())
	users.AuthorizationRouteRegister(v1.Group("/auth"))

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
