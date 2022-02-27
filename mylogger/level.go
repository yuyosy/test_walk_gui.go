package mylogger

type LogLevel int

const (
	FATAL LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
)

var LevelStr = map[LogLevel]string{
	FATAL: "FATAL",
	DEBUG: "DEBUG",
	INFO:  "INFO",
	WARN:  "WARN",
	ERROR: "ERROR",
}
