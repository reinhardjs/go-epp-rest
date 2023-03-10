package error_types

import (
	"fmt"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
)

type EPPCommandError struct {
	Result response.Result
}

func (e *EPPCommandError) Error() string {
	return fmt.Sprintf("%v", e.Result)
}
