package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type hostPresenter struct{}

func NewHostPresenter() presenter.HostPresenter {
	return &hostPresenter{}
}

func (p *hostPresenter) Check(response []byte) (responseObject response.CheckHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter Check: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) Create(response []byte) (responseObject response.CreateHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter Create: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) Update(response []byte) (responseObject response.UpdateHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter Update: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) Delete(response []byte) (responseObject response.DeleteHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter Delete: xml.Unmarshal"))
	}

	return
}

func (p *hostPresenter) Info(response []byte) (responseObject response.InfoHostResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "HostPresenter Info: xml.Unmarshal"))
	}

	return
}
