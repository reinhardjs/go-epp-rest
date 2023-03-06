package presenter

import (
	"fmt"
	"strings"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type contactPresenter struct{}

func NewContactPresenter() presenter.ContactPresenter {
	return &contactPresenter{}
}

func (p *contactPresenter) CheckSuccess(ctx infrastructure.Context, obj response.CheckContactResponse) {
	var res string

	for _, element := range obj.ResultData.CheckDatas {
		notStr := ""
		if element.Id.AvailKey == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Contact %s, contact %savailable\n", element.Id.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")

	ctx.String(200, res)
}

func (p *contactPresenter) CreateSuccess(ctx infrastructure.Context, obj response.CreateContactResponse) {
	var res string

	res += fmt.Sprintf("ID %s\n", obj.ResultData.CreateData.Id)
	res += fmt.Sprintf("Create Date %s\n", obj.ResultData.CreateData.CreateDate)
	res = strings.TrimSuffix(res, "\n")

	ctx.String(200, res)
}

func (p *contactPresenter) UpdateSuccess(ctx infrastructure.Context, obj response.UpdateContactResponse) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
}

func (p *contactPresenter) DeleteSuccess(ctx infrastructure.Context, obj response.DeleteContactResponse) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
}

func (p *contactPresenter) InfoSuccess(ctx infrastructure.Context, obj response.InfoContactResponse) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
}
