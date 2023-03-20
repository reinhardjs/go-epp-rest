package main

import (
	"log"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/config"
	"gitlab.com/merekmu/go-epp-rest/internal/server"
)

func periodicFree(d time.Duration) {
	tick := time.Tick(d)
	for _ = range tick {
		runtime.GC()
		debug.FreeOSMemory()
	}
}

func main() {
	go periodicFree(30 * time.Second)

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
