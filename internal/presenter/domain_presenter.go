package presenter

import (
	"fmt"
	"strings"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type domainPresenter struct{}

func NewDomainPresenter() presenter.DomainPresenter {
	return &domainPresenter{}
}

func (p *domainPresenter) Check(responseObject response.CheckDomainResponse) (res string) {

	for _, element := range responseObject.ResultData.CheckDatas {
		notStr := ""
		if element.Name.AvailKey == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Domain %s, domain %savailable\n", element.Name.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")

	return
}

func (p *domainPresenter) Create(responseObject response.CreateDomainResponse) (res string) {

	res += fmt.Sprintf("Name %s\n", responseObject.ResultData.CreatedData.Name)
	res += fmt.Sprintf("Create Date %s\n", responseObject.ResultData.CreatedData.CreatedDate)
	res += fmt.Sprintf("Expire Date %s\n", responseObject.ResultData.CreatedData.ExpiredDate)
	res = strings.TrimSuffix(res, "\n")

	return
}

func (p *domainPresenter) Delete(responseObject response.DeleteDomainResponse) (res string) {
	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)
	return
}

func (p *domainPresenter) Info(responseObject response.InfoDomainResponse) (res string) {
	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)
	return
}

func (p *domainPresenter) SecDNSUpdate(responseObject response.SecDNSUpdateResponse) (res string) {
	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)
	return
}
