package repository

type RegistrarRepository interface {
	Check(data interface{}) (string, error)
}
