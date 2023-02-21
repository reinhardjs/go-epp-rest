package interactor

type RegistrarInteractor interface {
	Check(data interface{}, ext string, langTag string) (string, error)
}
