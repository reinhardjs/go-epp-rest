package adapter

type EppClient interface {
	Send(data []byte) ([]byte, error)
	InitLogin(username, password string) ([]byte, error)
}
