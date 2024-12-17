package logbalancer

import (
	"context"
	"net/http"

	"connectrpc.com/connect"
	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs/pb_logsconnect"
)

type connectHandlers struct {
	lb *LogBalancer
}

func (c *connectHandlers) Send(ctx context.Context, c2 *connect.Request[pb_logs.RuntimeLogs]) (*connect.Response[pb_logs.SendResponse], error) {
	err := c.lb.HandleLog(c2.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&pb_logs.SendResponse{}), nil
}

func (c *connectHandlers) BatchSend(ctx context.Context, c2 *connect.Request[pb_logs.BatchSendRequest]) (*connect.Response[pb_logs.BatchSendResponse], error) {
	for _, l := range c2.Msg.Logs {
		c.lb.HandleLog(l)
	}
	return connect.NewResponse(&pb_logs.BatchSendResponse{}), nil
}

func HandleConnect(lb *LogBalancer) http.Handler {
	service := &connectHandlers{lb: lb}

	mux := http.NewServeMux()
	mux.Handle(pb_logsconnect.NewRpcLogsHandler(service))
	return mux
}
