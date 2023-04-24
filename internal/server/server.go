package server

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/config"
	"gitlab.com/merekmu/go-epp-rest/constants"
	"gitlab.com/merekmu/go-epp-rest/internal/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/adapter/mapper"
	"gitlab.com/merekmu/go-epp-rest/internal/delivery/http/router"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure/mysql"
	"gitlab.com/merekmu/go-epp-rest/internal/infrastructure/registry"
	"gitlab.com/merekmu/go-epp-rest/internal/utils"
)

type server struct {
	config      *config.Config
	sessionPool *utils.SessionPool
}

func NewServer(config *config.Config) *server {
	return &server{config: config}
}

func (s *server) Run() error {
	username := os.Getenv(constants.PAY_WEB_CC_REGISTRY_LOGIN_USERNAME)
	password := os.Getenv(constants.PAY_WEB_CC_REGISTRY_LOGIN_PASSWORD)
	tcpHost := os.Getenv(constants.PAY_WEB_CC_REGISTRY_TCP_HOST)
	apiHost := os.Getenv(constants.API_HOST)
	apiPort := os.Getenv(constants.API_PORT)

	tcpPort, err := strconv.Atoi(os.Getenv(constants.PAY_WEB_CC_REGISTRY_TCP_PORT))
	if err != nil {
		return errors.Wrap(err, "server run: strconv Atoi tcp port value from env")
	}

	tcpConfig := utils.TcpConfig{
		Host:         tcpHost,
		Port:         tcpPort,
		TLSCert:      s.config.PayWebCCCert,
		MaxOpenConn:  1,
		MaxIdleConns: 1,
	}
	tcpConnPool, err := utils.CreateTcpConnPool(&tcpConfig)
	if err != nil {
		return errors.Wrap(err, "server run: session pool create tcp conn pool")
	}
	logger := utils.GetLoggerInstance()
	eppClient := adapter.NewEppClient(tcpConnPool, logger)
	tcpConnPool.SetEppClient(eppClient)
	response, err := eppClient.InitLogin(username, password)
	if err != nil {
		log.Println(errors.Wrap(err, "server Run: eppClient.Login"))
		os.Exit(1)
	}

	mysqlConn, err := mysql.NewMysqlConnection(s.config.Mysql)
	if err != nil {
		return errors.Wrap(err, "server Run: mysql.NewMysqlConnection()")
	}

	xmlMapper := mapper.NewXMLMapper()
	dtoToEntityMapper := mapper.NewDtoToEntityMapper()
	registry := registry.NewRegistry(eppClient, mysqlConn, xmlMapper, dtoToEntityMapper)
	router := router.NewRouter(registry.NewAppController())

	log.Println("Login command result :")
	log.Println(string(response))

	router.Run(fmt.Sprintf("%v:%v", apiHost, apiPort))

	return nil
}
