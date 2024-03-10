package middleware

import (
	"xgo/core"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// RequestID 颁发RequestID
// TODO 应该由网关颁发
func RequestID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := ctx.Request.Header.Get(core.RequestIDKey)
		if requestId == "" {
			requestId = strings.Replace(uuid.New().String(), "-", "", -1)
		}
		ctx.Set(core.RequestIDKey, requestId)
	}
}
