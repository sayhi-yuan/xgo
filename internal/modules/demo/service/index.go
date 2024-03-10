package demo

import (
	"xgo/core"
	"fmt"
)

func (s service) Index(ctx *core.Context) {
	fmt.Print(ctx.UserInfo)
}
