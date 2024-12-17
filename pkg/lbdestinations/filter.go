package lbdestinations

import (
	"errors"

	"github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"
)

type SeverityFilter struct {
	UNSPECIFIED func(incomingLog *pb_logs.RuntimeLogs) error
	DEBUG       func(incomingLog *pb_logs.RuntimeLogs) error
	INFO        func(incomingLog *pb_logs.RuntimeLogs) error
	WARN        func(incomingLog *pb_logs.RuntimeLogs) error
	ERROR       func(incomingLog *pb_logs.RuntimeLogs) error
	CRITICAL    func(incomingLog *pb_logs.RuntimeLogs) error
	SUCCESS     func(incomingLog *pb_logs.RuntimeLogs) error
}

func FilterBySeverity(filter SeverityFilter) func(incomingLog *pb_logs.RuntimeLogs) error {
	return func(incomingLog *pb_logs.RuntimeLogs) error {
		switch incomingLog.Severity {
		case pb_logs.Severity_SEVERITY_UNSPECIFIED:
			return filter.UNSPECIFIED(incomingLog)
		case pb_logs.Severity_SEVERITY_DEBUG:
			return filter.DEBUG(incomingLog)
		case pb_logs.Severity_SEVERITY_INFO:
			return filter.INFO(incomingLog)
		case pb_logs.Severity_SEVERITY_WARN:
			return filter.WARN(incomingLog)
		case pb_logs.Severity_SEVERITY_ERROR:
			return filter.ERROR(incomingLog)
		case pb_logs.Severity_SEVERITY_CRITICAL:
			return filter.CRITICAL(incomingLog)
		case pb_logs.Severity_SEVERITY_SUCCESS:
			return filter.SUCCESS(incomingLog)
		}
		return nil
	}
}

func Join(fn ...func(incomingLog *pb_logs.RuntimeLogs) error) func(incomingLog *pb_logs.RuntimeLogs) error {
	return func(incomingLog *pb_logs.RuntimeLogs) error {
		var err []error
		for _, f := range fn {
			err1 := f(incomingLog)
			if err1 != nil {
				err = append(err, err1)
			}
		}
		if len(err) > 0 {
			return errors.Join(err...)
		}
		return nil
	}
}
