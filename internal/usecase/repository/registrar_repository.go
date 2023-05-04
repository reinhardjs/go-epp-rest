package repository

type RegistrarRepository interface {
	SendCommand(data interface{}) ([]byte, error)
	SendCommandV2(input interface{}, output interface{}) error
}
