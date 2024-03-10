package core

import (
	"xgo/config"
	"context"
	"fmt"
	"strings"

	"git.qdreads.com/gotools/log"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func init() {
	initLog()
}

func initLog() {
	cfg := config.Cfg.Logger
	log.InitLog(log.Options{
		CtxFields:    []string{RequestIDKey},
		Filename:     fmt.Sprintf("%s/%s", cfg.Path, config.Cfg.Serve.Name),
		MaxCount:     cfg.MaxCount,
		CallerEnable: true,
	}, zap.String("serve", config.Cfg.Serve.Name))
}

// 日志上下文
var JobCtx = context.WithValue(context.Background(), RequestIDKey, (func() string {
	return strings.Replace(uuid.New().String(), "-", "", -1)
})())
