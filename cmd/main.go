package main

import (
	"log"
	"runtime"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/config"
	"gitlab.com/merekmu/go-epp-rest/internal/server"
)

func periodicFree(d time.Duration) {
	tick := time.Tick(d)
	for range tick {
		runtime.GC()
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
