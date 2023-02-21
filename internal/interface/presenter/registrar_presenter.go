package presenter

import (
	"encoding/xml"
	"log"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/interface/constraints"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type registrarPresenter[T constraints.RegistrarResponseConstraint] struct{}

func NewRegistrarPresenter[T constraints.RegistrarResponseConstraint]() presenter.RegistrarPresenter[T] {
	return &registrarPresenter[T]{}
}

func (p *registrarPresenter[T]) Check(response []byte) (responseObject T, err error) {

	if err := xml.Unmarshal(response, &responseObject); err != nil {
		log.Println(errors.Wrap(err, "Domain Controller: CheckDomain xml.Unmarshal"))
	}

	return
}
