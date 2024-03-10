package main

import (
	"context"

	"git.qdreads.com/gotools/corntab"
	"git.qdreads.com/gotools/log"

	"xgo/bootstrap"
)

func main() {
	corntab.New()
	bootstrap.JobRun()
	if err := corntab.Execute(); err != nil {
		log.Errorf(context.Background(), "定时任务启动失败, %+v", err)
	}
}
