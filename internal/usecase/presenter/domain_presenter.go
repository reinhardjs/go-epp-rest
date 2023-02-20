package presenter

type DomainPresenter interface {
	ResponseQueryDomain(response string) (string, error)
}
