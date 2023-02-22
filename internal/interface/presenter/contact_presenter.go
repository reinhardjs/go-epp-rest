package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type contactPresenter struct{}

func NewContactPresenter() presenter.ContactPresenter {
	return &contactPresenter{}
}

func (p *contactPresenter) MapCheckResponse(response []byte) (responseObject model.CheckContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter MapCheckResponse: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) MapCreateResponse(response []byte) (responseObject model.CreateContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter MapCreateResponse: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) MapUpdateResponse(response []byte) (responseObject model.UpdateContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter MapUpdateResponse: xml.Unmarshal"))
	}

	return
}

func (p *contactPresenter) MapDeleteResponse(response []byte) (responseObject model.DeleteContactResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "ContactPresenter MapDeleteResponse: xml.Unmarshal"))
	}

	return
}
