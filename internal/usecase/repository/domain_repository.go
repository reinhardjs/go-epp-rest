package repository

type DomainRepository interface {
	DoQueryDomain(domainList []string, ext string, langTag string) (string, error)
}
