package presenter

import (
	"fmt"
	"strings"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type hostPresenter struct{}

func NewHostPresenter() presenter.HostPresenter {
	return &hostPresenter{}
}

func (p *hostPresenter) CheckSuccess(ctx infrastructure.Context, responseObject response.CheckHostResponse) {
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
}

func (p *hostPresenter) CreateSuccess(ctx infrastructure.Context, responseObject response.CreateHostResponse) {
	var res string

	res += fmt.Sprintf("Name %s\n", responseObject.ResultData.CreateData.Name)
	res += fmt.Sprintf("Create Date %s\n", responseObject.ResultData.CreateData.CreateDate)
	res = strings.TrimSuffix(res, "\n")

	ctx.String(200, res)
}

func (p *hostPresenter) UpdateSuccess(ctx infrastructure.Context, responseObject response.UpdateHostResponse) {
	var res string

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
}

func (p *hostPresenter) DeleteSuccess(ctx infrastructure.Context, responseObject response.DeleteHostResponse) {
	var res string

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
}

func (p *hostPresenter) InfoSuccess(ctx infrastructure.Context, responseObject response.InfoHostResponse) {
	var res string

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
}

func (p *hostPresenter) CheckAndCreateSuccess(ctx infrastructure.Context, responseObject response.CreateHostResponse) {
	var res string

	res += fmt.Sprintf("Name %s\n", responseObject.ResultData.CreateData.Name)
	res += fmt.Sprintf("Create Date %s\n", responseObject.ResultData.CreateData.CreateDate)
	res = strings.TrimSuffix(res, "\n")

	ctx.String(200, res)
}
