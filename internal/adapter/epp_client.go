package adapter

import (
	"io"
	"log"
	"net"
	"os"
	"syscall"
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

	response, err = c.login(tcpConn.Conn, username, password)
	if err != nil {
		log.Println(errors.Wrap(err, "server Run: eppClient.Login"))
		os.Exit(1)
	}

	tcpConn.SetShouldLogin(false)

	c.sessionPool.Put(tcpConn)

	return response, nil
}

// Send will send data to the server.
func (c *eppClient) Send(data []byte) (response []byte, err error) {
	tcpConn, err := c.sessionPool.Get()
	if err != nil {
		return nil, errors.Wrap(err, "EppClient Send: c.connPool.Get")
	}

	if tcpConn.GetShouldLogin() == true {
		response, err = c.login(tcpConn.Conn, c.loginCred.username, c.loginCred.password)
		if err != nil {
			return nil, errors.Wrap(err, "server Run: eppClient.Login")
		}

		tcpConn.SetShouldLogin(false)
	}

	startTime := time.Now()
	defer func() {
		c.timeTrack(startTime, "epp command response")

		if c.isNetConnClosedErr(err) {
			tcpConn.SetShouldLogin(true)
			c.sessionPool.Throw()
			return
		}

		if tcpConn != nil {
			c.sessionPool.Put(tcpConn)
		}
	}()

	return c.write(tcpConn.Conn, data)
}

// login will perform a login to an EPP server.
func (c *eppClient) login(conn net.Conn, username, password string) ([]byte, error) {
	login := types.Login{
		ClientID: username,
		Password: password,
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
		err = errors.Wrap(err, "EppClient Send: registry_epp.WriteMessage")

		errConn := conn.Close()
		if errConn != nil {
			err = errors.Wrap(errConn, err.Error())
		}

		return nil, err
	}

	err = conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "EppClient Send: conn.SetReadDeadline")
	}

	msg, err := registry_epp.ReadMessage(conn)
	if err != nil {
		err = errors.Wrap(err, "EppClient Send: registry_epp.ReadMessage")

		errConn := conn.Close()
		if errConn != nil {
			err = errors.Wrap(errConn, err.Error())
		}

		return nil, err
	}

	return msg, nil
}

func (c *eppClient) timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func (c *eppClient) isNetConnClosedErr(err error) bool {
	err = errors.Cause(err)
	if err != nil {
		netErr, ok := err.(net.Error)
		if ok && netErr.Timeout() {
			return true
		} else {
			switch {
			case
				errors.Is(err, net.ErrClosed),
				errors.Is(err, io.EOF),
				errors.Is(err, syscall.EPIPE):
				return true
			default:
				return false
			}
		}
	}

	return false
}
