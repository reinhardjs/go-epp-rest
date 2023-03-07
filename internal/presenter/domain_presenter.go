package presenter

import (
	"fmt"
	"strings"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type domainPresenter struct{}

func NewDomainPresenter() presenter.DomainPresenter {
	return &domainPresenter{}
}

func (p *domainPresenter) CheckSuccess(ctx infrastructure.Context, obj response.CheckDomainResponse) {
	var res string

	for _, element := range obj.ResultData.CheckDatas {
		notStr := ""
		if element.Name.AvailKey == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Domain %s, domain %savailable\n", element.Name.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")

	ctx.String(200, res)
}

func (p *domainPresenter) CreateSuccess(ctx infrastructure.Context, obj response.CreateDomainResponse) {
	var res string

	res += fmt.Sprintf("Name %s\n", obj.ResultData.CreatedData.Name)
	res += fmt.Sprintf("Create Date %s\n", obj.ResultData.CreatedData.CreatedDate)
	res += fmt.Sprintf("Expire Date %s\n", obj.ResultData.CreatedData.ExpiredDate)
	res = strings.TrimSuffix(res, "\n")

	ctx.String(200, res)
}

func (p *domainPresenter) DeleteSuccess(ctx infrastructure.Context, obj response.DeleteDomainResponse) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
}

func (p *domainPresenter) InfoSuccess(ctx infrastructure.Context, obj response.InfoDomainResponse) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
}

func (p *domainPresenter) SecDNSUpdateSuccess(ctx infrastructure.Context, obj response.SecDNSUpdateResponse) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
}

func (p *domainPresenter) ContactUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
}
