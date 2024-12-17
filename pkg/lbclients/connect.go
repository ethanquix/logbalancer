package lbclients

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs/pb_logsconnect"
)

func NewConnectLBClient(addr string, password string) pb_logsconnect.RpcLogsClient {
	return pb_logsconnect.NewRpcLogsClient(http.DefaultClient, addr, connect.WithInterceptors(connect.UnaryInterceptorFunc(authInterceptor(password))))
}

func authInterceptor(password string) func(next connect.UnaryFunc) connect.UnaryFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			req.Header().Set("Authorization", password)
			return next(ctx, req)
		}
	}
}
