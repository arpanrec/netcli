package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var zapSugaredLogger *zap.SugaredLogger

func init() {
	debugMode := false
	osArgs := os.Args
	for _, arg := range osArgs {
		if arg == "--debug-logging" {
			debugMode = true
			break
		}
	}
	config := zap.NewProductionConfig()
	if debugMode {
		config = zap.NewDevelopmentConfig()
		config.DisableStacktrace = false
		config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.Development = true
	} else {
		config.DisableStacktrace = true
		config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
		config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		config.Development = false
	}
	logger, err := config.Build()

	if err != nil {
		log.Panicln("Failed to create logger", err)
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	zapSugaredLogger = logger.Sugar()
}
func Debug(args ...interface{}) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Debug(args...)
}
func Info(args ...interface{}) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Info(args...)
}
func Warn(args ...interface{}) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Warn(args...)
}
func Error(args ...interface{}) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Error(args...)
}
func DPanic(args ...interface{}) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).DPanic(args...)
}
func Panic(args ...interface{}) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Panic(args...)
}
func Fatal(args ...interface{}) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Fatal(args...)
}
