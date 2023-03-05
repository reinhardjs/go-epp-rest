package server

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/config"
	"gitlab.com/merekmu/go-epp-rest/constants"
	"gitlab.com/merekmu/go-epp-rest/internal/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/adapter/mapper"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/router"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure/mysql"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure/registry"
	"gitlab.com/merekmu/go-epp-rest/pkg/webcc_epp/utils"
)

type server struct {
	config      *config.Config
	sessionPool *utils.TcpConnPool
}

func NewServer(config *config.Config) *server {
	return &server{config: config}
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

	username := os.Getenv(constants.PAY_WEB_CC_REGISTRY_LOGIN_USERNAME)
	password := os.Getenv(constants.PAY_WEB_CC_REGISTRY_LOGIN_PASSWORD)
	tcpConfig := utils.TcpConfig{
		Host:    tcpHost,
		Port:    tcpPort,
		TLSCert: s.config.PayWebCCCert,
	}
	tcpConnPool, err := utils.CreateTcpConnPool(&tcpConfig)
	if err != nil {
		return errors.Wrap(err, "server run: session pool create tcp conn pool")
	}
	tcpConn, err := tcpConnPool.Get()
	if err != nil {
		return errors.Wrap(err, "server run: tcpConnPool get")
	}
	eppClient := adapter.NewEppClient(tcpConn.Conn)
	response, err := eppClient.Login(username, password)
	if err != nil {
		log.Println(errors.Wrap(err, "server Run: eppClient.Login"))
		os.Exit(1)
	}

	mysqlConn, err := mysql.NewMysqlConnection(s.config.Mysql)
	if err != nil {
		return errors.Wrap(err, "server Run: mysql.NewMysqlConnection()")
	}

	xmlMapper := mapper.NewXMLMapper()
	registry := registry.NewRegistry(eppClient, mysqlConn, xmlMapper)
	router := router.NewRouter(registry.NewAppController())

	log.Println("Login command result :")
	log.Println(string(response))

	router.Run("localhost:8080")

	return nil
}
