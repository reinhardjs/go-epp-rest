package main

import (
	"log"
	"runtime/debug"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/config"
	"gitlab.com/merekmu/go-epp-rest/internal/server"
)

func periodicFree(d time.Duration) {
	tick := time.Tick(d)
	for _ = range tick {
		debug.FreeOSMemory()
	}
}

func main() {
	go periodicFree(1 * time.Minute)

	cfg, err := config.InitConfig()

	if err != nil {
		log.Println(errors.Wrap(err, "main: init config"))
	}

	s := server.NewServer(cfg)
	err = s.Run()

	if err != nil {
		log.Println(errors.Wrap(err, "main: new server"))
	}
}
