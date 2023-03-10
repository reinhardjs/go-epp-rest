package error_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/error_types"
)

func ClientErrorHandler(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors.Last()
		cause := errors.Cause(err)

		switch cause.(type) {
		case *error_types.ControllerError:
			// TODO with Controller Error
		case *error_types.InteractorError:
			// TODO with Interactor Error
		case *error_types.PresenterError:
			// TODO with Presenter Error
		case *error_types.EPPCommandError:
			eppComandError := cause.(*error_types.EPPCommandError)
			c.String(200, "2400 Command failed; "+eppComandError.Result.Message)
		default:
			c.String(200, "2400 Command failed")
		}
	}
}
