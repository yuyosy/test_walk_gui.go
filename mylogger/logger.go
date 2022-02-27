package mylogger

import (
	"fmt"
	"log"
	"os"
)

type MyLogger struct {
	children []*log.Logger
}

func New(lgs ...*log.Logger) *MyLogger {
	lg := &MyLogger{
		children: lgs,
	}
	return lg
}

func (l *MyLogger) AddLogger(lgs ...*log.Logger) {
	l.children = append(l.children, lgs...)
}

func (l *MyLogger) Debug(v ...interface{})                 { l.Log(DEBUG, v...) }
func (l *MyLogger) Info(v ...interface{})                  { l.Log(INFO, v...) }
func (l *MyLogger) Warn(v ...interface{})                  { l.Log(WARN, v...) }
func (l *MyLogger) Error(v ...interface{})                 { l.Log(ERROR, v...) }
func (l *MyLogger) Print(v ...interface{})                 { l.LogPrint(v...) }
func (l *MyLogger) Println(v ...interface{})               { l.LogPrintln(v...) }
func (l *MyLogger) Printf(format string, v ...interface{}) { l.LogPrintf(format, v...) }
func (l *MyLogger) Fatal(v ...interface{})                 { l.LogFatal(v...) }
func (l *MyLogger) Fatalln(v ...interface{})               { l.LogFatalln(v...) }
func (l *MyLogger) Fatalf(format string, v ...interface{}) { l.LogFatalf(format, v...) }

func (l *MyLogger) Log(level LogLevel, v ...interface{}) {
	for _, lg := range l.children {
		lg.SetPrefix(LevelStr[level] + " ")
		lg.Output(2, fmt.Sprintln(v...))
	}
}

func (l *MyLogger) LogPrint(v ...interface{}) {
	for _, lg := range l.children {
		lg.Print(v...)
	}
}

func (l *MyLogger) LogPrintln(v ...interface{}) {
	for _, lg := range l.children {
		lg.Println(v...)
	}
}

func (l *MyLogger) LogPrintf(format string, v ...interface{}) {
	for _, lg := range l.children {
		lg.Printf(format, v...)
	}
}

func (l *MyLogger) LogFatal(v ...interface{}) {
	l.LogPrint(v...)
	os.Exit(1)
}

func (l *MyLogger) LogFatalln(v ...interface{}) {
	l.LogPrintln(v...)
	os.Exit(1)
}

func (l *MyLogger) LogFatalf(format string, v ...interface{}) {
	l.LogPrintf(format, v...)
	os.Exit(1)
}
