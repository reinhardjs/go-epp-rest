package presenter

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/error_types"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type hostPresenter struct{}

func NewHostPresenter() presenter.HostPresenter {
	return &hostPresenter{}
}

func (p *hostPresenter) Check(ctx infrastructure.Context, responseObject response.CheckHostResponse) (err error) {
	var resultCode = responseObject.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseObject.Result}, "HostPresenter Check: epp command error")
		return
	}

	var res string

	for _, element := range responseObject.ResultData.CheckDatas {
		notStr := ""
		if element.HostName.Available == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Host %s is %savailable\n", element.HostName.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")
	res = fmt.Sprintf("%v %v", responseObject.Result.Code, res)

	ctx.String(200, res)
	return
}

func (p *hostPresenter) Create(ctx infrastructure.Context, responseObject response.CreateHostResponse) (err error) {
	var resultCode = responseObject.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseObject.Result}, "HostPresenter Create: epp command error")
		return
	}

	var res string
	res += fmt.Sprintf("Name %s\n", responseObject.ResultData.CreateData.Name)
	res += fmt.Sprintf("Create Date %s\n", responseObject.ResultData.CreateData.CreateDate)
	res = strings.TrimSuffix(res, "\n")

	ctx.String(200, res)
	return
}

func (p *hostPresenter) Update(ctx infrastructure.Context, responseObject response.UpdateHostResponse) (err error) {
	var resultCode = responseObject.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseObject.Result}, "HostPresenter Update: epp command error")
		return
	}

	var res string
	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
	return
}

func (p *hostPresenter) Delete(ctx infrastructure.Context, responseObject response.DeleteHostResponse) (err error) {
	var resultCode = responseObject.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseObject.Result}, "HostPresenter Delete: epp command error")
		return
	}

	var res string
	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
	return
}

func (p *hostPresenter) Info(ctx infrastructure.Context, responseObject response.InfoHostResponse) (err error) {
	var resultCode = responseObject.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseObject.Result}, "HostPresenter Info: epp command error")
		return
	}

	var res string
	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
	return
}

func (p *hostPresenter) CheckAndCreate(ctx infrastructure.Context, responseObject response.CreateHostResponse) (err error) {
	var resultCode = responseObject.Result.Code
	if resultCode >= 2000 {
		err = errors.Wrap(&error_types.EPPCommandError{Result: responseObject.Result}, "HostPresenter CheckAndCreate: epp command error")
		return
	}

	var res string
	res += fmt.Sprintf("Name %s\n", responseObject.ResultData.CreateData.Name)
	res += fmt.Sprintf("Create Date %s\n", responseObject.ResultData.CreateData.CreateDate)
	res = strings.TrimSuffix(res, "\n")

	ctx.String(200, res)
	return
}
