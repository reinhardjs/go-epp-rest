package repository

type RegistrarRepository interface {
	Check(data interface{}) ([]byte, error)
}
