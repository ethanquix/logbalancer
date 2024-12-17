// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: pkg/model/pb_logs/logs.proto

package pb_logsconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// RpcLogsName is the fully-qualified name of the RpcLogs service.
	RpcLogsName = "logbalancer.logs.v1.RpcLogs"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// RpcLogsSendProcedure is the fully-qualified name of the RpcLogs's Send RPC.
	RpcLogsSendProcedure = "/logbalancer.logs.v1.RpcLogs/Send"
	// RpcLogsBatchSendProcedure is the fully-qualified name of the RpcLogs's BatchSend RPC.
	RpcLogsBatchSendProcedure = "/logbalancer.logs.v1.RpcLogs/BatchSend"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	rpcLogsServiceDescriptor         = pb_logs.File_pkg_model_pb_logs_logs_proto.Services().ByName("RpcLogs")
	rpcLogsSendMethodDescriptor      = rpcLogsServiceDescriptor.Methods().ByName("Send")
	rpcLogsBatchSendMethodDescriptor = rpcLogsServiceDescriptor.Methods().ByName("BatchSend")
)

// RpcLogsClient is a client for the logbalancer.logs.v1.RpcLogs service.
type RpcLogsClient interface {
	Send(context.Context, *connect.Request[pb_logs.RuntimeLogs]) (*connect.Response[pb_logs.SendResponse], error)
	BatchSend(context.Context, *connect.Request[pb_logs.BatchSendRequest]) (*connect.Response[pb_logs.BatchSendResponse], error)
}

// NewRpcLogsClient constructs a client for the logbalancer.logs.v1.RpcLogs service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRpcLogsClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) RpcLogsClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &rpcLogsClient{
		send: connect.NewClient[pb_logs.RuntimeLogs, pb_logs.SendResponse](
			httpClient,
			baseURL+RpcLogsSendProcedure,
			connect.WithSchema(rpcLogsSendMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		batchSend: connect.NewClient[pb_logs.BatchSendRequest, pb_logs.BatchSendResponse](
			httpClient,
			baseURL+RpcLogsBatchSendProcedure,
			connect.WithSchema(rpcLogsBatchSendMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// rpcLogsClient implements RpcLogsClient.
type rpcLogsClient struct {
	send      *connect.Client[pb_logs.RuntimeLogs, pb_logs.SendResponse]
	batchSend *connect.Client[pb_logs.BatchSendRequest, pb_logs.BatchSendResponse]
}

// Send calls logbalancer.logs.v1.RpcLogs.Send.
func (c *rpcLogsClient) Send(ctx context.Context, req *connect.Request[pb_logs.RuntimeLogs]) (*connect.Response[pb_logs.SendResponse], error) {
	return c.send.CallUnary(ctx, req)
}

// BatchSend calls logbalancer.logs.v1.RpcLogs.BatchSend.
func (c *rpcLogsClient) BatchSend(ctx context.Context, req *connect.Request[pb_logs.BatchSendRequest]) (*connect.Response[pb_logs.BatchSendResponse], error) {
	return c.batchSend.CallUnary(ctx, req)
}

// RpcLogsHandler is an implementation of the logbalancer.logs.v1.RpcLogs service.
type RpcLogsHandler interface {
	Send(context.Context, *connect.Request[pb_logs.RuntimeLogs]) (*connect.Response[pb_logs.SendResponse], error)
	BatchSend(context.Context, *connect.Request[pb_logs.BatchSendRequest]) (*connect.Response[pb_logs.BatchSendResponse], error)
}

// NewRpcLogsHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRpcLogsHandler(svc RpcLogsHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	rpcLogsSendHandler := connect.NewUnaryHandler(
		RpcLogsSendProcedure,
		svc.Send,
		connect.WithSchema(rpcLogsSendMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	rpcLogsBatchSendHandler := connect.NewUnaryHandler(
		RpcLogsBatchSendProcedure,
		svc.BatchSend,
		connect.WithSchema(rpcLogsBatchSendMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/logbalancer.logs.v1.RpcLogs/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case RpcLogsSendProcedure:
			rpcLogsSendHandler.ServeHTTP(w, r)
		case RpcLogsBatchSendProcedure:
			rpcLogsBatchSendHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedRpcLogsHandler returns CodeUnimplemented from all methods.
type UnimplementedRpcLogsHandler struct{}

func (UnimplementedRpcLogsHandler) Send(context.Context, *connect.Request[pb_logs.RuntimeLogs]) (*connect.Response[pb_logs.SendResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("logbalancer.logs.v1.RpcLogs.Send is not implemented"))
}

func (UnimplementedRpcLogsHandler) BatchSend(context.Context, *connect.Request[pb_logs.BatchSendRequest]) (*connect.Response[pb_logs.BatchSendResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("logbalancer.logs.v1.RpcLogs.BatchSend is not implemented"))
}
