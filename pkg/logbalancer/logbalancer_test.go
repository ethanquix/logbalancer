package logbalancer

import (
	"context"
	"testing"
	"time"

	"connectrpc.com/connect"
	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
	"github.com/ethanquix/logbalancer/pkg/lbclients"
	"github.com/ethanquix/logbalancer/pkg/lbdestinations"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNew(t *testing.T) {
	// telegram
	//tg, err := lbdestinations.NewTelegram("")
	//require.NoError(t, err)
	//
	//// slack
	//sl := lbdestinations.NewSlack("")

	lb := New(WithPassword("password"), WithPort(8011))
	//lb.On("/", tg.SendTo(6665765455))
	//lb.On("/", sl.SendTo("jarvis"))
	lb.On("/", lbdestinations.StdoutSend)
	// pretend connect
	go func() {
		time.Sleep(1 * time.Second)
		cc := lbclients.NewConnectLBClient("http://localhost:8011/connect", "password")
		_, err := cc.Send(context.Background(), connect.NewRequest(&pb_logs.RuntimeLogs{
			LogDate:  timestamppb.New(time.Now()),
			Severity: pb_logs.Severity_SEVERITY_INFO,
			Source:   "test",
			Message:  "Hello from test",
			Context:  nil,
			Path:     "/",
		}))
		require.NoError(t, err)
	}()

	require.NoError(t, lb.Run())
}
