package adapter

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
	"gitlab.com/merekmu/go-epp-rest/pkg/webcc_epp/utils"
)

// Client represents an EPP client.
type eppClient struct {
	// connPool holds the TCP connections to the server.
	connPool *utils.TcpConnPool
}

func NewEppClient(connPool *utils.TcpConnPool) adapter.EppClient {
	return &eppClient{
		connPool: connPool,
	}
}

func (c *eppClient) timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

// Send will send data to the server.
func (c *eppClient) Send(data []byte) ([]byte, error) {
	tcpConn, err := c.connPool.Get()

	defer c.timeTrack(time.Now(), "epp command response")

	if err != nil {
		return nil, errors.Wrap(err, "EppClient Send: c.connPool.Get")
	}

	err = registry_epp.WriteMessage(tcpConn.Conn, data)
	if err != nil {
		_ = tcpConn.Conn.Close()

		return nil, errors.Wrap(err, "EppClient Send: registry_epp.WriteMessage")
	}

	err = tcpConn.Conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	if err != nil {
		return nil, errors.Wrap(err, "EppClient Send: tcpConn.Conn.SetReadDeadline")
	}

	msg, err := registry_epp.ReadMessage(tcpConn.Conn)
	if err != nil {
		_ = tcpConn.Conn.Close()

		return nil, errors.Wrap(err, "EppClient Send: registry_epp.ReadMessage")
	}

	c.connPool.Put(tcpConn)

	return msg, nil
}

// Login will perform a login to an EPP server.
func (c *eppClient) Login(username, password string) ([]byte, error) {
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

	return c.Send(encoded)
}
