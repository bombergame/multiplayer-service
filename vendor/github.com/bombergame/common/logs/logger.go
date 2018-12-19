package logs

import (
	"fmt"
	"runtime/debug"

	"github.com/sirupsen/logrus"
)

//Logger is the wrapper for a logging library
type Logger struct {
	logger *logrus.Logger
}

//NewLogger creates logger instance
func NewLogger() *Logger {
	return &Logger{
		logger: logrus.StandardLogger(),
	}
}

//Info writes info level log
func (log *Logger) Info(args ...interface{}) {
	log.logger.Info(args...)
}

//Error writes error level log
func (log *Logger) Error(args ...interface{}) {
	log.withStackTrace(log.logger.Error, args)
}

//Fatal writes fatal level log
func (log *Logger) Fatal(args ...interface{}) {
	log.withStackTrace(log.logger.Fatal, args)
}

//LogrusLogger returns logrus.Logger object
func (log *Logger) LogrusLogger() *logrus.Logger {
	return log.logger
}

func (log *Logger) withStackTrace(f func(...interface{}), args ...interface{}) {
	f(append(args, fmt.Sprintf("Stack trace:\n%s", string(debug.Stack()))))
}
