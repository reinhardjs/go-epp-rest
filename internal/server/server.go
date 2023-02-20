package server

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/config"
	"gitlab.com/merekmu/go-epp-rest/internal/constants"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/pkg/webcc_epp/utils"
)

type server struct {
	cfg         *config.Config
	sessionPool *utils.TcpConnPool
}

func NewServer(cfg *config.Config) *server {
	return &server{cfg: cfg}
}

func (s *server) Run() error {
	err := godotenv.Load()
	if err != nil {
		return errors.Wrap(err, "server run: godotenv load")
	}

	tcpHost := os.Getenv(constants.PAY_WEB_CC_REGISTRY_TCP_HOST)
	tcpPort, err := strconv.Atoi(os.Getenv(constants.PAY_WEB_CC_REGISTRY_TCP_PORT))

	if err != nil {
		return errors.Wrap(err, "server run: strconv Atoi tcp port value from env")
	}

	tcpConfig := utils.TcpConfig{
		Host:    tcpHost,
		Port:    tcpPort,
		TLSCert: s.cfg.PayWebCCCert,
	}
	tcpConnPool, err := utils.CreateTcpConnPool(&tcpConfig)
	if err != nil {
		return errors.Wrap(err, "server run: session pool create tcp conn pool")
	}

	tcpConn, err := tcpConnPool.Get()
	if err != nil {
		return errors.Wrap(err, "server run: tcpConnPool get")
	}

	username := os.Getenv(constants.PAY_WEB_CC_REGISTRY_LOGIN_USERNAME)
	password := os.Getenv(constants.PAY_WEB_CC_REGISTRY_LOGIN_PASSWORD)
	eppClient := infrastructure.NewEppClient(tcpConn.Conn)

	response, err := eppClient.Login(username, password)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	log.Println("Login command result :")
	log.Println(string(response))

	return nil
}
