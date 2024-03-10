package demo

import (
	"github.com/gin-gonic/gin"
)

func Router(app *gin.Engine) {
	group := app.Group("demo")

	// demo
	{
		demo := newDemoHandle()
		group.POST("", demo.Index)
	}
}
