package infrastructure

import (
	"net"
	"time"

	"github.com/bombsimon/epp-go"
	"github.com/bombsimon/epp-go/types"
	"github.com/pkg/errors"
)

// Client represents an EPP client.
type eppClient struct {
	// conn holds the TCP connection to the server.
	conn net.Conn
}

type EppClient interface {
	Send(data []byte) ([]byte, error)
	Login(username, password string) ([]byte, error)
}

func NewEppClient(conn net.Conn) EppClient {
	return &eppClient{
		conn: conn,
	}
}

// Send will send data to the server.
func (c *eppClient) Send(data []byte) ([]byte, error) {
	err := epp.WriteMessage(c.conn, data)
	if err != nil {
		_ = c.conn.Close()

		return nil, errors.Wrap(err, "EppClient Send: epp.WriteMessage")
	}

	_ = c.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	msg, err := epp.ReadMessage(c.conn)
	if err != nil {
		_ = c.conn.Close()

		return nil, errors.Wrap(err, "EppClient Send: epp.ReadMessage")
	}

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

	encoded, err := epp.Encode(login, epp.ClientXMLAttributes())
	if err != nil {
		return nil, errors.Wrap(err, "EppClient Send: epp.ReadMessage")
	}

	return c.Send(encoded)
}
