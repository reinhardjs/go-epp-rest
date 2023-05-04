package utils

import (
	"fmt"
	"strings"
	"sync"

	"github.com/sirupsen/logrus"
)

var lock = &sync.Mutex{}

var instance *logger

type logger struct {
	logChannel chan string
	log        *logrus.Logger
}

type Logger interface {
	Info(args ...interface{})
}

func GetLoggerInstance() Logger {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()

		if instance == nil {
			instance = &logger{
				logChannel: make(chan string),
				log:        logrus.New(),
			}

			instance.Run()
		}

		return instance
	}

	return instance
}

func (l *logger) Run() {
	l.log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		TimestampFormat: "2006-01-02 15:04:05.999999999",
	})

	go func(logger *logger) {
		for msg := range logger.logChannel {
			// Create a file for writing logs
			// file, err := os.OpenFile("logs/api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

			// defer file.Close()

			// if err != nil {
			// 	panic(err)
			// }

			// Set the logger to write output to both the file and terminal
			// l.log.SetOutput(io.MultiWriter(os.Stdout))
			logger.log.Info(msg)
		}
	}(l)
}

func (l *logger) Info(args ...interface{}) {
	l.logChannel <- strings.Trim(fmt.Sprintf("%v", args), "[]")
}
