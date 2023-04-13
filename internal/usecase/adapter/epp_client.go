package adapter

import "net"

type EppClient interface {
	Send(data []byte) ([]byte, error)
	InitLogin(username, password string) ([]byte, error)
	DoLogin(conn net.Conn) ([]byte, error)
	SendHello(conn net.Conn) ([]byte, error)
}
