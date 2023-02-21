package presenter

type RegistrarPresenter interface {
	ResponseCheck(response string) (string, error)
}
