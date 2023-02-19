package main

import (
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/config"
	"gitlab.com/merekmu/go-epp-rest/internal/server"
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
