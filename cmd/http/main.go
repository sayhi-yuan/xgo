package main

import (
	"xgo/bootstrap"
	"xgo/config"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.New()

	app.Use(static.Serve("/static", static.LocalFile("./static", false)))
	// 加载项目
	bootstrap.Run(app)

	// 启动服务
	serve := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Cfg.Serve.Port),
		Handler: app,
	}
	go func() {
		if err := serve.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic("服务启动失败")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := serve.Shutdown(ctx); err != nil {
		// TODO 关闭服务异常
	}
}
