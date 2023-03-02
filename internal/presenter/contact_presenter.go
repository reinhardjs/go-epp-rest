package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/adapter/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type contactPresenter struct{}

func NewContactPresenter() presenter.ContactPresenter {
	return &contactPresenter{}
}

func (p *contactPresenter) Check(response []byte) (responseObject response.CheckContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter Check: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) Create(response []byte) (responseObject response.CreateContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter Create: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) Update(response []byte) (responseObject response.UpdateContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter Update: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) Delete(response []byte) (responseObject response.DeleteContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter Delete: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) Info(response []byte) (responseObject response.InfoContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter Info: xml.Unmarshal"))
	}

	return
}
