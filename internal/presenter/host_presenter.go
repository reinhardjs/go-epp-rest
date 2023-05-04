package presenter

import (
	"fmt"
	"strings"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/utils"
)

type hostPresenter struct{}

func NewHostPresenter() presenter.HostPresenter {
	return &hostPresenter{}
}

func (p *hostPresenter) Check(ctx infrastructure.Context, responseObject response.CheckHostResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	for _, element := range responseObject.ResultData.CheckDatas {
		notStr := ""
		if element.HostName.Available == 0 {
			notStr = "not "
		}
		buffer.WriteString(fmt.Sprintf("Host %s is %savailable\n", element.HostName.Value, notStr))
	}

	res := strings.TrimSuffix(buffer.String(), "\n")
	res = fmt.Sprintf("%v %v", responseObject.Result.Code, res)
	ctx.String(200, res)
	return
}

func (p *hostPresenter) Create(ctx infrastructure.Context, responseObject response.CreateHostResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	buffer.WriteString(fmt.Sprintf("%d %s", responseObject.Result.Code, responseObject.Result.Message))

	ctx.String(200, buffer.String())
	return
}

func (p *hostPresenter) Update(ctx infrastructure.Context, responseObject response.UpdateHostResponse) (err error) {
	var res string
	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
	return
}

func (p *hostPresenter) Delete(ctx infrastructure.Context, responseObject response.DeleteHostResponse) (err error) {
	var res string
	res = fmt.Sprintf("%v %v", responseObject.Result.Code, responseObject.Result.Message)

	ctx.String(200, res)
	return
}

func (p *hostPresenter) Info(ctx infrastructure.Context, responseObject response.InfoHostResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	buffer.WriteString(fmt.Sprintf("1000 Host[%s]\n", responseObject.ResultData.InfoData.Name))
	buffer.WriteString(fmt.Sprintf("ClientID[%s]\n", responseObject.ResultData.InfoData.ClientID))
	buffer.WriteString(fmt.Sprintf("Created[%s]\n", responseObject.ResultData.InfoData.CreateDate.Local().Format("2006-01-02 15:04:05")))
	buffer.WriteString(fmt.Sprintf("UpID[%s]\n", responseObject.ResultData.InfoData.UpdateID))
	buffer.WriteString(fmt.Sprintf("Updated[%s]\n", responseObject.ResultData.InfoData.UpdateDate.Local().Format("2006-01-02 15:04:05")))

	for _, address := range responseObject.ResultData.InfoData.Address {
		buffer.WriteString(fmt.Sprintf("IPAddress[%s]Type[%s]", address.Address, address.IPType))
	}

	ctx.String(200, buffer.String())
	return
}

func (p *hostPresenter) CheckAndCreate(ctx infrastructure.Context, responseObject response.CreateHostResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	buffer.WriteString(fmt.Sprintf("%d %s", responseObject.Result.Code, responseObject.Result.Message))

	ctx.String(200, buffer.String())
	return
}
