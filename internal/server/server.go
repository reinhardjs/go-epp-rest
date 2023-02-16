package server

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gitlab.com/reinhardjs/go-epp-rest/internal/config"
	"gitlab.com/reinhardjs/go-epp-rest/internal/constants"
	"gitlab.com/reinhardjs/go-epp-rest/pkg/webcc_epp"
	"gitlab.com/reinhardjs/go-epp-rest/pkg/webcc_epp/session_pool"
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
		return errors.Wrap(err, "server run: godotenv load")
	}

	tcpConfig := session_pool.TcpConfig{
		Host:    os.Getenv(constants.PAY_WEB_CC_REGISTRY_TCP_HOST),
		Port:    1700,
		TLSCert: s.cfg.PayWebCCCert,
	}
	tcpConnPool, err := session_pool.CreateTcpConnPool(&tcpConfig)
	if err != nil {
		return errors.Wrap(err, "server run: session pool create tcp conn pool")
	}

	tcpConn, err := tcpConnPool.Get()
	if err != nil {
		return errors.Wrap(err, "server run: tcpConnPool get")
	}

	username := os.Getenv(constants.PAY_WEB_CC_REGISTRY_LOGIN_USERNAME)
	password := os.Getenv(constants.PAY_WEB_CC_REGISTRY_LOGIN_PASSWORD)
	eppClient := webcc_epp.NewClient(tcpConn.Conn)

	response, err := eppClient.Login(username, password)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	log.Println("Login command result :")
	log.Println(string(response))

	return nil
}
