package adapter

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
	"gitlab.com/merekmu/go-epp-rest/pkg/webcc_epp/utils"
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
}

func NewEppClient(connPool *utils.SessionPool) adapter.EppClient {
	return &eppClient{
		sessionPool: connPool,
	}
}

func (c *eppClient) InitLogin(username string, password string) (response []byte, err error) {
	c.loginCred = loginCred{username, password}

	tcpConn, err := c.sessionPool.Get()
	if err != nil {
		return nil, errors.Wrap(err, "EppClient Send: c.connPool.Get")
	}

	response, err = c.DoLogin(tcpConn.Conn)
	if err != nil {
		log.Println(errors.Wrap(err, "server Run: eppClient.Login"))
		os.Exit(1)
	}

	c.sessionPool.Put(tcpConn)

	return response, nil
}

// Send will send data to the server.
func (c *eppClient) Send(data []byte) (response []byte, err error) {
	var session *utils.Session

	session, err = c.sessionPool.Get()
	defer c.sessionPool.Put(session)

	if err != nil {
		return nil, errors.Wrap(err, "EppClient Send: c.connPool.Get")
	}

	var startTime time.Time

	var tcpConn net.Conn = session.GetTcpConn()
	if tcpConn != nil {
		startTime = time.Now()
		response, err = c.write(tcpConn, data)
		c.trackTime(startTime, "epp command response")
	}

	if tcpConn == nil || c.isNetConnClosedErr(err) {
		tcpConn, err = c.sessionPool.RenewTcpConn(session)
		if err != nil {
			return
		}

		startTime = time.Now()
		response, err = c.write(tcpConn, data)
		c.trackTime(startTime, "epp command response")
	}

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
		return nil, errors.Wrap(err, "EppClient Send: registry_epp.ReadMessage")
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
	log.Printf("%s took %s", name, elapsed)
}

func (c *eppClient) isNetConnClosedErr(err error) bool {
	err = errors.Cause(err)
	if err != nil {
		return true
	}

	return false
}
