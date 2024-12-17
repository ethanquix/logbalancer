package utils

import "github.com/ethanquix/logbalancer/gen/go/pkg/model/pb_logs"

func SeverityToString(s pb_logs.Severity) string {
	switch s {
	case pb_logs.Severity_SEVERITY_UNSPECIFIED:
		return "Unspecified"
	case pb_logs.Severity_SEVERITY_DEBUG:
		return "Debug"
	case pb_logs.Severity_SEVERITY_INFO:
		return "Info"
	case pb_logs.Severity_SEVERITY_WARN:
		return "Warn"
	case pb_logs.Severity_SEVERITY_ERROR:
		return "Error"
	case pb_logs.Severity_SEVERITY_CRITICAL:
		return "Critical"
	case pb_logs.Severity_SEVERITY_SUCCESS:
		return "Success"
	}
	return "Unspecified"
}
