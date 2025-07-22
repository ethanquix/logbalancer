package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
	"github.com/ethanquix/logbalancer/pkg/lbclients"
	"github.com/ethanquix/logbalancer/pkg/logbalancer"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var testsMap = sync.Map{}

func init() {
	testsMap.Store("hello_world", false)
}

func TestBalancer(t *testing.T) {
	msgReceivedRoot := atomic.Int64{}
	lb := logbalancer.New(logbalancer.WithPort("7222"), logbalancer.WithPassword("password"))

	// == ROUTES ==

	// Hello world
	lb.On("/hello_world", func(incomingLog *pb_logs.RuntimeLogs) error {
		require.Equal(t, "hello world", incomingLog.Message)
		testsMap.Store("hello_world", true)
		return nil
	})
	// All
	lb.On("/", func(incomingLog *pb_logs.RuntimeLogs) error {
		msgReceivedRoot.Add(1)
		return nil
	})

	// == SEND

	go func() {
		err := lb.Run()
		if err != nil && err.Error() != "http: Server closed" {
			panic(err)
		}
	}()
	time.Sleep(time.Millisecond * 100) // give time for the server to start

	// send logs
	client := lbclients.NewConnectLBClient("http://localhost:7222/connect", "password")

	// Hello World
	_, err := client.Send(t.Context(), connect.NewRequest(&pb_logs.RuntimeLogs{
		Severity: pb_logs.Severity_SEVERITY_INFO,
		Source:   "test",
		Message:  "hello world",
		Path:     "/hello_world",
		LogDate:  timestamppb.New(time.Now()),
	}))
	require.NoError(t, err)

	// Check all tests
	time.Sleep(time.Second)

	testsMap.Range(func(key, value interface{}) bool {
		if value == false {
			t.Errorf("test %s never ran", key)
		}
		return true
	})

	require.Equal(t, int64(1), msgReceivedRoot.Load())

	lb.Stop()
}
