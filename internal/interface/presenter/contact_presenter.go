package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/responses"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type contactPresenter struct{}

func NewContactPresenter() presenter.ContactPresenter {
	return &contactPresenter{}
}

func (p *contactPresenter) MapCheckResponse(response []byte) (responseObject responses.CheckContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter MapCheckResponse: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) MapCreateResponse(response []byte) (responseObject responses.CreateContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter MapCreateResponse: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) MapUpdateResponse(response []byte) (responseObject responses.UpdateContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter MapUpdateResponse: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) MapDeleteResponse(response []byte) (responseObject responses.DeleteContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter MapDeleteResponse: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) MapInfoResponse(response []byte) (responseObject responses.InfoContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter MapInfoResponse: xml.Unmarshal"))
	}

	return
}
