package gocommon

type LogLevel int32

const (
	LogLevel_TRACE LogLevel = 0
	LogLevel_DEBUG LogLevel = 1
	LogLevel_INFO  LogLevel = 2
	LogLevel_WARN  LogLevel = 3
	LogLevel_ERROR LogLevel = 4
	LogLevel_OFF   LogLevel = 5
)

func (l LogLevel) String() string {
	switch l {
	case LogLevel_TRACE:
		return "TRACE"
	case LogLevel_DEBUG:
		return "DEBUG"
	case LogLevel_INFO:
		return "INFO"
	case LogLevel_WARN:
		return "WARN"
	case LogLevel_ERROR:
		return "ERROR"
	case LogLevel_OFF:
		return "OFF"
	default:
		return "UNKNOWN"
	}
}
