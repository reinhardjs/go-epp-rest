package presenter

import (
	"fmt"
	"strings"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type contactPresenter struct{}

func NewContactPresenter() presenter.ContactPresenter {
	return &contactPresenter{}
}

func (p *contactPresenter) Check(responseObject response.CheckContactResponse) (res string) {

	for _, element := range responseObject.ResultData.CheckDatas {
		notStr := ""
		if element.Id.AvailKey == 0 {
			notStr = "not "
		}
		res += fmt.Sprintf("Contact %s, contact %savailable\n", element.Id.Value, notStr)
	}
	res = strings.TrimSuffix(res, "\n")

	return
}

func (p *contactPresenter) Create(responseObject response.CreateContactResponse) (res string) {

	res += fmt.Sprintf("ID %s\n", responseObject.ResultData.CreateData.Id)
	res += fmt.Sprintf("Create Date %s\n", responseObject.ResultData.CreateData.CreateDate)
	res = strings.TrimSuffix(res, "\n")

	return
}

func (p *contactPresenter) Update(responseObject response.UpdateContactResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}

func (p *contactPresenter) Delete(responseObject response.DeleteContactResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}

func (p *contactPresenter) Info(responseObject response.InfoContactResponse) (res string) {

	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	return
}
