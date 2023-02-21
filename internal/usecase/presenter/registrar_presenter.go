package presenter

type RegistrarPresenter interface {
	ResponseCheck(response []byte) (string, error)
}
