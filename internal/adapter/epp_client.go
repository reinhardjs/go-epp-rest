package adapter

import (
	"net"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

// Client represents an EPP client.
type eppClient struct {
	// conn holds the TCP connection to the server.
	conn net.Conn
}

func NewEppClient(conn net.Conn) adapter.EppClient {
	return &eppClient{
		conn: conn,
	}
}

// Send will send data to the server.
func (c *eppClient) Send(data []byte) ([]byte, error) {
	err := registry_epp.WriteMessage(c.conn, data)
	if err != nil {
		_ = c.conn.Close()

		return nil, errors.Wrap(err, "EppClient Send: registry_epp.WriteMessage")
	}

	_ = c.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	msg, err := registry_epp.ReadMessage(c.conn)
	if err != nil {
		_ = c.conn.Close()

		return nil, errors.Wrap(err, "EppClient Send: registry_epp.ReadMessage")
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

	encoded, err := registry_epp.Encode(login, registry_epp.ClientXMLAttributes())
	if err != nil {
		return nil, errors.Wrap(err, "EppClient Send: registry_epp.ReadMessage")
	}

	return c.Send(encoded)
}
