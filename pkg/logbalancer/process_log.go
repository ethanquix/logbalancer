package logbalancer

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
)

func (lb *LogBalancer) HandleLog(incomingLog *pb_logs.RuntimeLogs) error {
	// match with route
	if incomingLog == nil {
		return fmt.Errorf("log is nil")
	}
	if incomingLog.Path == "" {
		return fmt.Errorf("log path is empty")
	}
	if incomingLog.LogDate == nil {
		return fmt.Errorf("log date is nil")
	}
	if incomingLog.Message == "" {
		return fmt.Errorf("message is empty")
	}
	if incomingLog.Severity == pb_logs.Severity_SEVERITY_UNSPECIFIED {
		return fmt.Errorf("severity is unspecified")
	}
	if incomingLog.Source == "" {
		return fmt.Errorf("source is empty")
	}

	for path, targets := range lb.listeners {
		for _, t := range targets {
			if _, isMatch := t.path.Match(incomingLog.Path); isMatch {
				err := t.fn(incomingLog)
				if err != nil {
					log.Errorf("sending log for path %s: %v", path, err)
				}
			}
		}
	}
	return nil
}
