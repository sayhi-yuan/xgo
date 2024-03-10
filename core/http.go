package core

import (
	"context"

	"git.qdreads.com/gotools/log"
	"github.com/guonaihong/gout"
	"github.com/guonaihong/gout/dataflow"
)

// 参考文档https://github.com/guonaihong/gout

func POST(ctx context.Context, url string) *dataflow.DataFlow {
	return gout.POST(url).Debug(goutDebug(ctx))
}

func GET(ctx context.Context, url string) *dataflow.DataFlow {
	return gout.GET(url).Debug(goutDebug(ctx))
}

// 添加debug
func goutDebug(ctx context.Context) gout.DebugOpt {
	return gout.DebugFunc(func(o *gout.DebugOption) {
		o.Debug = true
		o.Write = log.SafeWriter(ctx)
	})
}
