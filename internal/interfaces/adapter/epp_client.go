package adapter

type EppClient interface {
	Send(data []byte) ([]byte, error)
	Login(username, password string) ([]byte, error)
}
