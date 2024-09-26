package logger

import (
	"github.com/dobyte/due/v2/log"
)

func init() {
	log.SetLogger(log.NewLogger(log.WithCallerSkip(2)))
}

// Debug 打印调试日志
func Debug(a ...interface{}) {
	log.GetLogger().Debug(a...)
}

// Debugf 打印调试模板日志
func Debugf(format string, a ...interface{}) {
	log.GetLogger().Debugf(format, a...)
}

// Info 打印信息日志
func Info(a ...interface{}) {
	log.GetLogger().Info(a...)
}

// Infof 打印信息模板日志
func Infof(format string, a ...interface{}) {
	log.GetLogger().Infof(format, a...)
}

// Warn 打印警告日志
func Warn(a ...interface{}) {
	log.GetLogger().Warn(a...)
}

// Warnf 打印警告模板日志
func Warnf(format string, a ...interface{}) {
	log.GetLogger().Warnf(format, a...)
}

// Error 打印错误日志
func Error(a ...interface{}) {
	log.GetLogger().Error(a...)
}

// Errorf 打印错误模板日志
func Errorf(format string, a ...interface{}) {
	log.GetLogger().Errorf(format, a...)
}

// Fatal 打印致命错误日志
func Fatal(a ...interface{}) {
	log.GetLogger().Fatal(a...)
}

// Fatalf 打印致命错误模板日志
func Fatalf(format string, a ...interface{}) {
	log.GetLogger().Fatalf(format, a...)
}

// Panic 打印Panic日志
func Panic(a ...interface{}) {
	log.GetLogger().Panic(a...)
}

// Panicf 打印Panic模板日志
func Panicf(format string, a ...interface{}) {
	log.Panicf(format, a...)
}
