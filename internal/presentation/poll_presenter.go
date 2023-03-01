package presentation

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/model/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type pollPresenter struct{}

func NewPollPresenter() presenter.PollPresenter {
	return &pollPresenter{}
}

func (p *pollPresenter) Acknowledge(response []byte) (responseObject response.PollAckResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "PollPresenter Acknowledge: xml.Unmarshal"))
	}

	return
}

func (p *pollPresenter) Request(response []byte) (responseObject response.PollRequestResponse, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "PollPresenter Request: xml.Unmarshal"))
	}

	return
}
