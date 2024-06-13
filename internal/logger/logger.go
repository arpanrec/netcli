package logger

import (
	"log"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapSugaredLogger *zap.SugaredLogger

func SetUpLogger() {
	debugMode := false
	osArgs := os.Args
	for _, arg := range osArgs {
		if arg == "--debug-logging" {
			debugMode = true
			break
		}
	}
	if !debugMode {
		allEnv := os.Environ()
		for _, env := range allEnv {
			if strings.HasPrefix(env, "DEBUG=") {
				kv := strings.Split(env, "=")
				dM, errDm := strconv.ParseBool(kv[1])
				if errDm != nil {
					log.Panic("failed to parse DEBUG env variable, " + errDm.Error())
				}
				debugMode = dM
			}
		}
	}

	config := zap.NewProductionConfig()
	if debugMode {
		config = zap.NewDevelopmentConfig()
		config.DisableStacktrace = false
		config.Development = true
	} else {
		config.DisableStacktrace = true
		config.Development = false
	}
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.EncoderConfig.EncodeName = zapcore.FullNameEncoder
	config.EncoderConfig.TimeKey = "timestamp"
	logger, err := config.Build()

	if err != nil {
		log.Panicln("Failed to create logger", err)
	}
	defer func(logger *zap.Logger) {
		loggerSyncErr := logger.Sync()
		if loggerSyncErr != nil {
			log.Println("Warn:: Failed to sync logger", loggerSyncErr)
		}
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
