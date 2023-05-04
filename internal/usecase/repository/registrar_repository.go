package repository

type RegistrarRepository interface {
	SendCommand(data interface{}) ([]byte, error)
	SendCommandV2(data interface{}) (interface{}, error)
}
