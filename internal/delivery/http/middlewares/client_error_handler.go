package middlewares

import (
	"bytes"
	"fmt"
	"regexp"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/error_types"
	"gitlab.com/merekmu/go-epp-rest/internal/utils"
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

		logger := utils.GetLoggerInstance()
		logger.Info(err)

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
			var buffer bytes.Buffer
			eppCommandErr := cause.(*error_types.EPPCommandError)
			resultCode := registry_epp.ResultCode(eppCommandErr.Result.Code)

			buffer.WriteString(fmt.Sprintf("%d %s", resultCode.Code(), resultCode.Message()))

			eppCommandError := cause.(*error_types.EPPCommandError)

			if len(eppCommandError.Result.Value.Texts) > 0 {
				value := eppCommandError.Result.Value.Texts[0]
				re := regexp.MustCompile(`\[(.*?)\]`)
				match := re.FindStringSubmatch(value)
				if len(match) > 1 {
					buffer.WriteString(fmt.Sprintf(" | %s", match[1]))
				}
			}

			if eppCommandError.Result.ExternalValue != nil {
				buffer.WriteString(fmt.Sprintf(" | %s %s", eppCommandError.Result.ExternalValue.Value.ReasonCode, eppCommandError.Result.ExternalValue.Reason))
			}

			c.String(200, buffer.String())
		default:
			c.String(200, "2400 Command failed")
		}
	}
}
