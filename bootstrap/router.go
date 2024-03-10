package bootstrap

import (
	"xgo/internal/modules/demo"

	"github.com/gin-gonic/gin"
)

func notAuthRouter(app *gin.Engine) {

}

func loadRouter(app *gin.Engine) {
	demo.Router(app)
}
