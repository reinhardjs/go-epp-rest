package middlewares

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/error_types"
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp"
)

func ClientErrorHandler(c *gin.Context) {
	c.Next()

	defer func() {
		runtime.GC()
	}()

	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err
		cause := errors.Cause(err)

		switch cause.(type) {
		case *error_types.RequestTimeOutError:
			c.String(408, "2400 Command failed; Request time out")
		case *error_types.ControllerError:
			// TODO with Controller Error
		case *error_types.InteractorError:
			// TODO with Interactor Error
		case *error_types.PresenterError:
			// TODO with Presenter Error
		case *error_types.EPPCommandError:
			// TODO with EPPCommand Error
			eppCommandErr := cause.(*error_types.EPPCommandError)
			resultCode := registry_epp.ResultCode(eppCommandErr.Result.Code)
			c.String(200, fmt.Sprintf("%d %s", resultCode.Code(), resultCode.Message()))
		default:
			c.String(200, "2400 Command failed")
		}
	}
}
