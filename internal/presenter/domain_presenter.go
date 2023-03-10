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

func (p *domainPresenter) CheckSuccess(ctx infrastructure.Context, obj response.CheckDomainResponse) (err error) {
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
	return
}

func (p *domainPresenter) CreateSuccess(ctx infrastructure.Context, obj response.CreateDomainResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %s", obj.Result.Code, obj.ResultData.CreatedData.ExpiredDate)

	ctx.String(200, res)
	return
}

func (p *domainPresenter) DeleteSuccess(ctx infrastructure.Context, obj response.DeleteDomainResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
	return
}

func (p *domainPresenter) InfoSuccess(ctx infrastructure.Context, obj response.InfoDomainResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
	return
}

func (p *domainPresenter) SecDNSUpdateSuccess(ctx infrastructure.Context, obj response.SecDNSUpdateResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
	return
}

func (p *domainPresenter) ContactUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
	return
}

func (p *domainPresenter) StatusUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
	return
}

func (p *domainPresenter) AuthInfoUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
	return
}

func (p *domainPresenter) NameserverUpdateSuccess(ctx infrastructure.Context, obj response.DomainUpdateResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
	return
}

func (p *domainPresenter) RenewSuccess(ctx infrastructure.Context, obj response.RenewDomainResponse) (err error) {
	var res string

	if obj.Result.Code == 1000 {
		res = fmt.Sprintf("%v %v", obj.Result.Code, obj.ResultData.RenewedData.ExpiredDate)
	} else {
		res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)
	}

	ctx.String(200, res)
	return
}
