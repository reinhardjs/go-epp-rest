package repository

type RegistrarRepository interface {
	Check(list []string, ext string, langTag string) (string, error)
}
