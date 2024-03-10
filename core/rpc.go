package core

import (
	"xgo/config"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func init() {
	initRpc()
}

var RpcConn *grpc.ClientConn

func initRpc() {
	adds := fmt.Sprintf("%s:%s", config.Cfg.FileService.Host, config.Cfg.FileService.Port)
	conn, err := grpc.Dial(adds, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	RpcConn = conn
}

const ServerNameKey = "server-name"

func RpcSwap(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, ServerNameKey, config.Cfg.Serve.Name)
}
