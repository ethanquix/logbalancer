package lbdestinations

import (
	"github.com/charmbracelet/log"
	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
)

func StdoutSend(incomingLog *pb_logs.RuntimeLogs) error {
	log.Infof("received: %v", incomingLog)
	return nil
}
