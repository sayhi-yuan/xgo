package bootstrap

import (
	"xgo/internal/middleware"
	"github.com/gin-gonic/gin"
)

func loadMiddleware(app *gin.Engine) {
	app.Use(
		middleware.Auth(),
	)
}
