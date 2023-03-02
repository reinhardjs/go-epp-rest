package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/adapter/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type domainPresenter struct{}

func NewDomainPresenter() presenter.DomainPresenter {
	return &domainPresenter{}
}

func (p *domainPresenter) Check(response []byte) (responseObject response.CheckDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainPresenter Check: xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) Create(response []byte) (responseObject response.CreateDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainPresenter Create: xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) Delete(response []byte) (responseObject response.DeleteDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainPresenter Delete: xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) Info(response []byte) (responseObject response.InfoDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainPresenter Info: xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) SecDNSUpdate(response []byte) (responseObject response.SecDNSUpdateResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainPresenter MapSecDNSUpdateResponse: xml.Unmarshal"))
	}

	return
}
