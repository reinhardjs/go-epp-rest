package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/responses"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type hostPresenter struct{}

func NewHostPresenter() presenter.HostPresenter {
	return &hostPresenter{}
}

func (p *hostPresenter) MapCheckResponse(response []byte) (responseObject responses.CheckHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter MapCheckResponse: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) MapCreateResponse(response []byte) (responseObject responses.CreateHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter MapCreateResponse: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) MapUpdateResponse(response []byte) (responseObject responses.UpdateHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter MapUpdateResponse: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) MapDeleteResponse(response []byte) (responseObject responses.DeleteHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter MapDeleteResponse: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) MapInfoResponse(response []byte) (responseObject responses.InfoHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter MapInfoResponse: xml.Unmarshal"))
	}

	return
}
