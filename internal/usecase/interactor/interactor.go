package interactor

import (
	"gitlab.com/merekmu/go-epp-rest/internal/interface/constraints"
)

type RegistrarInteractor[T constraints.RegistrarResponseConstraint] interface {
	Send(data interface{}) (T, error)
	Check(data interface{}, ext string, langTag string) (res string, err error)
}
