package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type hostPresenter struct{}

func NewHostPresenter() presenter.HostPresenter {
	return &hostPresenter{}
}

func (p *hostPresenter) MapCheckResponse(response []byte) (responseObject model.CheckHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter MapCheckResponse: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) MapCreateResponse(response []byte) (responseObject model.CreateHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter MapCreateResponse: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) MapUpdateResponse(response []byte) (responseObject model.UpdateHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter MapUpdateResponse: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) MapDeleteResponse(response []byte) (responseObject model.DeleteHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter MapDeleteResponse: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) MapInfoResponse(response []byte) (responseObject model.InfoHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter MapInfoResponse: xml.Unmarshal"))
	}

	return
}
