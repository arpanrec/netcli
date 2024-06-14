package logger

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapSugaredLogger *zap.SugaredLogger
var zapLogger *zap.Logger
var debugMode = false

func setDebugMode() {
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
				value := strings.Replace(env, "DEBUG=", "", 1)
				dM, errDm := strconv.ParseBool(value)
				if errDm != nil {
					debugMode = false
				}
				debugMode = dM
				break
			}
		}
	}
}

func setUpZapLogger() {
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
	zapLogger = logger

	if err != nil {
		log.Panicln("Failed to create logger", err)
	}

	zapSugaredLogger = zapLogger.Sugar()
}

func SetUpLogger() {
	setDebugMode()
	setUpZapLogger()
}

func Debug(v ...any) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Debug(v...)
}

func Info(v ...any) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Info(v...)
}

func Warn(v ...any) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Warn(v...)
}

func Error(v ...any) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Error(v...)
}

// Panic is equivalent to [Print] followed by a call to panic().
func Panic(v ...any) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Panic(v...)
	panic(fmt.Sprintln(v...))
}

// Fatal is equivalent to [Print] followed by a call to [os.Exit](1).
func Fatal(v ...any) {
	zapSugaredLogger.WithOptions(zap.AddCallerSkip(1)).Fatal(v...)
	os.Exit(1)
}

func Sync() {
	err := zapLogger.Sync()
	if err != nil {
		return
		// log.Println("WARN:: Failed to sync logger", err) // TODO: https://github.com/uber-go/zap/issues/328
	}
}
