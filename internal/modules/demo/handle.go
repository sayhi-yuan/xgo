package demo

import (
	"xgo/core"
	"xgo/internal/modules/demo/dto"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Interface interface {
	Index(ctx *gin.Context)
}

type demoHandle struct {
	core.BaseHandle
}

func newDemoHandle() Interface {
	return &demoHandle{}
}

func (handle demoHandle) Index(ctx *gin.Context) {
	param := dto.DemoRequest{}
	if err := handle.BindParam(ctx, &param); err != nil {
		handle.FailError(ctx, err)
		return
	}

	fmt.Println("Hello World")

	var res any
	core.POST(handle.Swap(ctx), "tapi.xayabx.com/auth/login").SetJSON(map[string]any{
		"email": "111@qq.com",
	}).BindJSON(&res).Do()
	fmt.Println(res)
}
