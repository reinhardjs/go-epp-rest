package webcc_epp

import (
	"net"
	"time"

	"github.com/bombsimon/epp-go"
	"github.com/bombsimon/epp-go/types"
)

// Client represents an EPP client.
type Client struct {
	// conn holds the TCP connection to the server.
	conn net.Conn
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

// Send will send data to the server.
func (c *Client) Send(data []byte) ([]byte, error) {
	err := epp.WriteMessage(c.conn, data)
	if err != nil {
		_ = c.conn.Close()

		return nil, err
	}

	_ = c.conn.SetReadDeadline(time.Now().Add(2 * time.Second))
	msg, err := epp.ReadMessage(c.conn)
	if err != nil {
		_ = c.conn.Close()

		return nil, err
	}

	return msg, nil
}

// Login will perform a login to an EPP server.
func (c *Client) Login(username, password string) ([]byte, error) {
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
		return nil, err
	}

	return c.Send(encoded)
}
