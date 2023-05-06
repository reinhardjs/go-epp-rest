package utils

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var loggerLock = &sync.Mutex{}

var loggerInstance *logger

type logger struct {
	logChannel chan string
	log        *zap.Logger
}

type Logger interface {
	Info(args ...interface{})
}

func GetLoggerInstance() Logger {
	loggerLock.Lock()
	defer loggerLock.Unlock()

	if loggerInstance == nil {
		if loggerInstance == nil {
			loggerInstance = &logger{}

			loggerInstance.Init()
		}

		return loggerInstance
	}

	return loggerInstance
}

func (l *logger) Init() {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	// fileEncoder := zapcore.NewConsoleEncoder(config)
	// logFile, _ := os.OpenFile("logs/api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// writer := zapcore.AddSync(logFile)

	consoleEncoder := zapcore.NewConsoleEncoder(config)
	defaultLogLevel := zapcore.DebugLevel

	core := zapcore.NewTee(
		// zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	l.log = zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
	l.logChannel = make(chan string)

	go func(logger *logger) {
		for msg := range logger.logChannel {
			logger.log.Info(msg)
		}
	}(l)
}

func (l *logger) Info(args ...interface{}) {
	l.logChannel <- strings.Trim(fmt.Sprintf("%v", args), "[]")
}
