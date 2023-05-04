package utils

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var loggerLock = &sync.Mutex{}

var loggerInstance *logger

type logger struct {
	logChannel chan string
	log        *logrus.Logger
}

type Logger interface {
	Info(args ...interface{})
}

func GetLoggerInstance() Logger {
	loggerLock.Lock()
	defer loggerLock.Unlock()

	if loggerInstance == nil {
		if loggerInstance == nil {
			loggerInstance = &logger{
				logChannel: make(chan string),
				log:        logrus.New(),
			}

			loggerInstance.Run()
		}

		return loggerInstance
	}

	return loggerInstance
}

func (l *logger) Run() {
	// Create a file for writing logs
	file, err := os.OpenFile("logs/api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	// Set the logger to write output to both the file and terminal
	l.log.SetOutput(io.MultiWriter(os.Stdout, file))

	l.log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		TimestampFormat: "2006-01-02 15:04:05.999999999",
	})

	go func(logger *logger) {
		for msg := range logger.logChannel {
			logger.log.Info(msg)
		}
	}(l)
}

func (l *logger) Info(args ...interface{}) {
	l.logChannel <- strings.Trim(fmt.Sprintf("%v", args), "[]")
}
