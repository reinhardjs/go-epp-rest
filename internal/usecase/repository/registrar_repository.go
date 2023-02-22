package repository

type RegistrarRepository interface {
	SendCommand(data interface{}) ([]byte, error)
}
