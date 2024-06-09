package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

var zapLog *zap.SugaredLogger

func init() {
	config := zap.NewDevelopmentConfig()
	config.DisableStacktrace = false
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build()

	if err != nil {
		log.Panicln("Failed to create logger", err)
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	zapLog = logger.Sugar()
}
func Debug(args ...interface{}) {
	zapLog.WithOptions(zap.AddCallerSkip(1)).Debug(args...)
}
func Info(args ...interface{}) {
	zapLog.WithOptions(zap.AddCallerSkip(1)).Info(args...)
}
func Warn(args ...interface{}) {
	zapLog.WithOptions(zap.AddCallerSkip(1)).Warn(args...)
}
func Error(args ...interface{}) {
	zapLog.WithOptions(zap.AddCallerSkip(1)).Error(args...)
}
func DPanic(args ...interface{}) {
	zapLog.WithOptions(zap.AddCallerSkip(1)).DPanic(args...)
}
func Panic(args ...interface{}) {
	zapLog.WithOptions(zap.AddCallerSkip(1)).Panic(args...)
}
func Fatal(args ...interface{}) {
	zapLog.WithOptions(zap.AddCallerSkip(1)).Fatal(args...)
}
