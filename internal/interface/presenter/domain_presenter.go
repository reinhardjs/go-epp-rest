package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/responses"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type domainPresenter struct{}

func NewDomainPresenter() presenter.DomainPresenter {
	return &domainPresenter{}
}

func (p *domainPresenter) MapCheckResponse(response []byte) (responseObject responses.CheckDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainPresenter MapCheckResponse: xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) MapCreateResponse(response []byte) (responseObject responses.CreateDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainPresenter MapCreateResponse: xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) MapDeleteResponse(response []byte) (responseObject responses.DeleteDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainPresenter MapDeleteResponse: xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) MapInfoResponse(response []byte) (responseObject responses.InfoDomainResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainPresenter MapInfoResponse: xml.Unmarshal"))
	}

	return
}

func (p *domainPresenter) MapSecDNSUpdateResponse(response []byte) (responseObject responses.SecDNSUpdateResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "DomainPresenter MapSecDNSUpdateResponse: xml.Unmarshal"))
	}

	return
}
