package server

import (
	"crypto/tls"
	"log"
	"os"

	"github.com/bombsimon/epp-go"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gitlab.com/reinhardjs/go-epp-rest/internal/config"
	"gitlab.com/reinhardjs/go-epp-rest/internal/constants"
	"gitlab.com/reinhardjs/go-epp-rest/internal/session_pool"
)

type server struct {
	cfg         *config.Config
	sessionPool *session_pool.TcpConnPool
}

func NewServer(cfg *config.Config) *server {
	return &server{cfg: cfg}
}

func (s *server) Run() error {
	err := godotenv.Load()

	if err != nil {
		return errors.Wrap(err, "server run: godotenv load:")
	}

	client := &epp.Client{
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
			Certificates:       []tls.Certificate{*s.cfg.Cert},
		},
	}

	greeting, err := client.Connect(os.Getenv(constants.PAY_WEB_CC_REGISTRY_TCP_HOST))

	if err != nil {
		log.Fatal(errors.Wrap(err, "server run: tcp client connect"))
	}

	log.Println(greeting)

	return nil
}
