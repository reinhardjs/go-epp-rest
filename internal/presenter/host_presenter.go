package presenter

import (
	"fmt"
	"strings"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type hostPresenter struct{}

func NewHostPresenter() presenter.HostPresenter {
	return &hostPresenter{}
}

func (p *hostPresenter) Check(responseObject response.CheckHostResponse) (res string) {

	for _, element := range responseObject.ResultData.CheckDatas {
		notStr := ""
		if element.HostName.AvailKey == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Host %s, host %savailable\n", element.HostName.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")

	return
}

func (p *hostPresenter) Create(responseObject response.CreateHostResponse) (res string) {

	res += fmt.Sprintf("Name %s\n", responseObject.ResultData.CreateData.Name)
	res += fmt.Sprintf("Create Date %s\n", responseObject.ResultData.CreateData.CreateDate)
	res = strings.TrimSuffix(res, "\n")

	return
}

func (p *hostPresenter) Update(responseObject response.UpdateHostResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}

func (p *hostPresenter) Delete(responseObject response.DeleteHostResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}

func (p *hostPresenter) Info(responseObject response.InfoHostResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}
