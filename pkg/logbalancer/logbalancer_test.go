package logbalancer

import (
	"context"
	"fmt"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
	"github.com/ethanquix/logbalancer/pkg/lbclients"
	"github.com/ethanquix/logbalancer/pkg/lbdestinations"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNew(t *testing.T) {
	t.Skip()
	// telegram
	//tg, err := lbdestinations.NewTelegram("")
	//require.NoError(t, err)
	//
	//// slack
	sl, _ := lbdestinations.NewSlack("")

	lb := New(WithPassword("password"), WithPort("8011"))
	//lb.On("/", tg.SendTo(6665765455))
	//lb.On("/", sl.SendTo("jarvis"))
	lb.On("/", func(incomingLog *pb_logs.RuntimeLogs) error {
		js, _ := protojson.Marshal(incomingLog)
		fmt.Println(string(js))
		return nil
	})

	lb.On("/", lbdestinations.FilterBySeverity(lbdestinations.SeverityFilter{
		WARN:  sl.SendTo("warning"),
		ERROR: sl.SendTo("error"),
		DEBUG: lbdestinations.Join(sl.SendTo("debug"), lbdestinations.StdoutSend),
	}))
	// pretend connect
	go func() {
		time.Sleep(1 * time.Second)
		cc := lbclients.NewConnectLBClient("http://localhost:8011/connect", "password")
		_, err := cc.Send(context.Background(), connect.NewRequest(&pb_logs.RuntimeLogs{
			LogDate:  timestamppb.New(time.Now()),
			Severity: pb_logs.Severity_SEVERITY_INFO,
			Source:   "test",
			Message:  "Hello from test",
			Context: map[string]string{
				"client": "123",
			},
			Tags: map[string]string{
				"version": "456",
			},
			Path: "/",
		}))
		require.NoError(t, err)
	}()

	require.NoError(t, lb.Run())
}
