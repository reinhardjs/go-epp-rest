package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type domainPresenter struct{}

func NewDomainPresenter() presenter.DomainPresenter {
	return &domainPresenter{}
}

func (p *domainPresenter) MapCheckResponse(response []byte) (responseObject model.CheckDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "Domain Controller: CheckDomain xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) MapCreateResponse(response []byte) (responseObject model.CreateDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "Domain Controller: CreateDomain xml.Unmarshal"))
	}

	return
}
