package constraints

import (
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
)

type RegistrarResponseConstraint interface {
	model.CheckDomainResponse | model.CheckContactResponse
}
