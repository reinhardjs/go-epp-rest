package presenter

import (
	"gitlab.com/merekmu/go-epp-rest/internal/interface/constraints"
)

type RegistrarPresenter[T constraints.RegistrarResponseConstraint] interface {
	ResponseCheck(response []byte) (string, error)
}
