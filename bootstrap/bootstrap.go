package bootstrap

import (
	_ "xgo/config"
	_ "xgo/core"
	"xgo/internal/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Run(app *gin.Engine) {
	// 健康检测
	app.GET("health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "success")
	})

	// 全局中间件
	app.Use(middleware.RequestID())

	notAuthRouter(app)

	loadMiddleware(app)
	// 加载路由
	loadRouter(app)
}
