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
		log.Println(errors.Wrap(err, "DomainController CheckDomain: xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) MapCreateResponse(response []byte) (responseObject model.CreateDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainController CreateDomain: xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) MapDeleteResponse(response []byte) (responseObject model.DeleteDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainController DeleteDomain: xml.Unmarshal"))
	}

	return
}
