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
			if filter.UNSPECIFIED != nil {
				return filter.UNSPECIFIED(incomingLog)
			}
		case pb_logs.Severity_SEVERITY_DEBUG:
			if filter.DEBUG != nil {
				return filter.DEBUG(incomingLog)
			}
		case pb_logs.Severity_SEVERITY_INFO:
			if filter.INFO != nil {
				return filter.INFO(incomingLog)
			}
		case pb_logs.Severity_SEVERITY_WARN:
			if filter.WARN != nil {
				return filter.WARN(incomingLog)
			}
		case pb_logs.Severity_SEVERITY_ERROR:
			if filter.ERROR != nil {
				return filter.ERROR(incomingLog)
			}
		case pb_logs.Severity_SEVERITY_CRITICAL:
			if filter.CRITICAL != nil {
				return filter.CRITICAL(incomingLog)
			}
		case pb_logs.Severity_SEVERITY_SUCCESS:
			if filter.SUCCESS != nil {
				return filter.SUCCESS(incomingLog)
			}
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
