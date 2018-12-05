package logs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime/debug"
)

type Logger struct {
	logger *logrus.Logger
}

func NewLogger() *Logger {
	return &Logger{
		logger: logrus.StandardLogger(),
	}
}

func (log *Logger) Info(args ...interface{}) {
	log.logger.Info(args...)
}

func (log *Logger) Error(args ...interface{}) {
	log.withStackTrace(log.logger.Error, args)
}

func (log *Logger) Fatal(args ...interface{}) {
	log.withStackTrace(log.logger.Fatal, args)
}

func (log *Logger) AsLogrusLogger() *logrus.Logger {
	return log.logger
}

func (log *Logger) withStackTrace(f func(...interface{}), args ...interface{}) {
	f(append(args, fmt.Sprintf("Stack trace:\n%s", string(debug.Stack()))))
}
