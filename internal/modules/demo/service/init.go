package demo

import "xgo/core"

type Interface interface {
	Index(ctx *core.Context)
}

func NewDemoService() Interface {
	return &service{}
}

type service struct {
}
