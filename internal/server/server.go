package server

import (
	"encoding/xml"
	"log"
	"os"
	"strconv"

	"github.com/bombsimon/epp-go"
	"github.com/bombsimon/epp-go/types"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/config"
	"gitlab.com/merekmu/go-epp-rest/constants"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
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

	domainCheck := types.DomainCheckType{
		Check: types.DomainCheck{
			Names: []string{"reinhard.com", "jonathan.com"},
		},
	}

	encoded, err := epp.Encode(domainCheck, epp.ClientXMLAttributes())
	if err != nil {
		log.Println("Error :", err)
	}

	byteResponse, err := eppClient.Send(encoded)
	if err != nil {
		log.Println("Error :", err)
	}

	log.Println("Response :", string(byteResponse))

	responseObj := model.DomainCheckResponse{}

	if err := xml.Unmarshal(byteResponse, &responseObj); err != nil {
		log.Println("Error :", err)
	}

	log.Println("Response Obj :", responseObj)

	return nil
}
