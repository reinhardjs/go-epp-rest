package adapter

import (
	"bytes"
	"fmt"
	"net"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter"
	"gitlab.com/merekmu/go-epp-rest/internal/utils"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

type loginCred struct {
	username string
	password string
}

// Client represents an EPP client.
type eppClient struct {
	// sessionPool holds the TCP connections to the server.
	sessionPool *utils.SessionPool
	loginCred   loginCred
	generator   utils.IDGenerator
	logger      utils.Logger
}

func NewEppClient(connPool *utils.SessionPool, logger utils.Logger, username string, password string) adapter.EppClient {
	return &eppClient{
		sessionPool: connPool,
		loginCred:   loginCred{username, password},
		generator:   utils.NewGenerator(),
		logger:      logger,
	}
}

// Send will send data to the server.
func (c *eppClient) Send(data []byte) (response []byte, err error) {
	var buffer bytes.Buffer
	var session *utils.Session

	session, err = c.sessionPool.Get()
	defer c.sessionPool.Put(session)

	requestId := c.generator.GenerateRequestId()
	sessionId := session.Id

	buffer.WriteString(fmt.Sprintln("\n"+requestId, " | ", sessionId))
	buffer.WriteString(fmt.Sprintf("%v%v", " --------------- XML Request: --------------- \n", string(data)))
	c.logger.Info(buffer.String())

	if err != nil {
		return nil, errors.Wrap(err, "EppClient Send: c.connPool.Get")
	}

	var startTime time.Time
	var tcpConn net.Conn = session.GetTcpConn()
	startTime = time.Now()
	if tcpConn != nil {
		response, err = c.write(tcpConn, data)
	}

	if tcpConn == nil || c.isNetConnClosedErr(err) {
		tcpConn, err = c.sessionPool.RenewTcpConn(session)
		if err != nil {
			return
		}

		response, err = c.write(tcpConn, data)
	}

	buffer.Reset()
	buffer.WriteString(fmt.Sprintln("\n"+requestId, " | ", sessionId))
	buffer.WriteString(fmt.Sprintf("%v%v", " --------------- XML Response: ---------------", string(response)))
	c.logger.Info(buffer.String())
	c.trackTime(startTime, "epp command response\n\n\n")

	return
}

// Login will perform a Login to an EPP server.
func (c *eppClient) DoLogin(conn net.Conn) ([]byte, error) {
	login := types.Login{
		ClientID: c.loginCred.username,
		Password: c.loginCred.password,
		Options: types.LoginOptions{
			Version:  "1.0",
			Language: "en",
		},
		Services: types.LoginServices{
			ObjectURI: []string{
				"urn:ietf:params:xml:ns:domain-1.0",
				"urn:ietf:params:xml:ns:contact-1.0",
				"urn:ietf:params:xml:ns:host-1.0",
			},
			ServiceExtension: types.LoginServiceExtension{
				ExtensionURI: []string{
					"urn:ietf:params:xml:ns:secDNS-1.0",
					"urn:ietf:params:xml:ns:secDNS-1.1",
				},
			},
		},
	}

	encoded, err := registry_epp.Encode(login, registry_epp.ClientXMLAttributes())
	if err != nil {
		return nil, errors.Wrap(err, "EppClient DoLogin: registry_epp.Encode")
	}

	return c.write(conn, encoded)
}

func (c *eppClient) SendHello(conn net.Conn) (response []byte, err error) {
	hello := types.Hello{}

	encoded, err := registry_epp.Encode(hello, registry_epp.ClientXMLAttributes())
	if err != nil {
		return nil, errors.Wrap(err, "EppClient SendHello: registry_epp.Encode")
	}

	return c.write(conn, encoded)
}

func (c *eppClient) write(conn net.Conn, data []byte) (response []byte, err error) {
	err = registry_epp.WriteMessage(conn, data)
	if err != nil {
		_ = conn.Close()

		return nil, errors.Wrap(err, "EppClient Send: registry_epp.WriteMessage")
	}

	err = conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		_ = conn.Close()

		return nil, errors.Wrap(err, "EppClient Send: conn.SetReadDeadline")
	}

	msg, err := registry_epp.ReadMessage(conn)
	if err != nil {
		_ = conn.Close()

		return nil, errors.Wrap(err, "EppClient Send: registry_epp.ReadMessage")
	}

	return msg, nil
}

func (c *eppClient) trackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	c.logger.Info(fmt.Sprintf("%s took %s", name, elapsed))
}

func (c *eppClient) isNetConnClosedErr(err error) bool {
	err = errors.Cause(err)
	if err != nil {
		return true
	}

	return false
}
