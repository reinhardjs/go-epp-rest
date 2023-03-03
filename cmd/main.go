package main

import (
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/config"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/server"
)

func main() {
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
