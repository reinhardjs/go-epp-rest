package presenter

import (
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/presenter/infrastructure"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
	"gitlab.com/merekmu/go-epp-rest/internal/utils"
)

type domainPresenter struct{}

func NewDomainPresenter() presenter.DomainPresenter {
	return &domainPresenter{}
}

func (p *domainPresenter) CheckSuccess(ctx infrastructure.Context, obj response.CheckDomainResponse) (err error) {
	var res string
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	for _, element := range obj.ResultData.CheckDatas {
		notStr := ""
		if element.Name.AvailKey == 0 {
			notStr = "not "
		}
		buffer.WriteString(fmt.Sprintf("Domain %s, domain %savailable\n", element.Name.Value, notStr))
	}

	res = strings.TrimSuffix(buffer.String(), "\n")
	ctx.String(200, res)
	return
}

func (p *domainPresenter) CreateSuccess(ctx infrastructure.Context, obj response.CreateDomainResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	layoutFormat := "2006-01-02T15:04:05.999999999Z"
	expiringDate, errParse := time.Parse(layoutFormat, obj.ResultData.CreatedData.ExpiredDate)
	if errParse != nil {
		err = errors.Wrap(errParse, "PollInteractor Poll: QueueDate time.Parse")
		return
	}
	expiringDate = expiringDate.Local()
	expiringDate = expiringDate.Add(-(time.Hour * 8))

	buffer.WriteString(fmt.Sprintf("%s %s", "1000", expiringDate.Format("2006-01-02 15:04:05")))

	ctx.String(200, buffer.String())
	return
}

func (p *domainPresenter) DeleteSuccess(ctx infrastructure.Context, obj response.DeleteDomainResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
	return
}

func (p *domainPresenter) InfoSuccess(ctx infrastructure.Context, obj response.InfoDomainResponse) (err error) {
	buffer := utils.GetBufferPoolInstance().Get()
	defer utils.GetBufferPoolInstance().Put(buffer)

	buffer.WriteString(fmt.Sprintf("%s domain[%s]\n", "1000", obj.ResultData.InfoData.Name))
	buffer.WriteString(fmt.Sprintf("domainns[%s]\n", strings.Join(obj.ResultData.InfoData.NameServer.HostObject, ",")))
	buffer.WriteString(fmt.Sprintf("host[%s]\n", strings.Join(obj.ResultData.InfoData.Host, ",")))
	buffer.WriteString(fmt.Sprintf("regid[%s]\n", obj.ResultData.InfoData.Registrant))

	contactMap := make(map[string]string)
	for _, contact := range obj.ResultData.InfoData.Contact {
		contactMap[contact.Type] = contact.Name
	}

	buffer.WriteString(fmt.Sprintf("admid[%s]\n", contactMap["admin"]))
	buffer.WriteString(fmt.Sprintf("tecid[%s]\n", contactMap["tech"]))
	buffer.WriteString(fmt.Sprintf("bilid[%s]\n", contactMap["billing"]))

	buffer.WriteString(fmt.Sprintf("authinfo[%s]\n", obj.ResultData.InfoData.AuthInfo.Password))

	createDate := *obj.ResultData.InfoData.CreateDate
	createDate = createDate.Local()
	createDate = createDate.Add(-(time.Hour * 8))
	buffer.WriteString(fmt.Sprintf("createdate[%s]\n", createDate.Format("2006-01-02 15:04:05")))

	expiryDate := *obj.ResultData.InfoData.ExpireDate
	expiryDate = expiryDate.Local()
	expiryDate = expiryDate.Add(-(time.Hour * 8))
	buffer.WriteString(fmt.Sprintf("expirydate[%s]\n", expiryDate.Format("2006-01-02 15:04:05")))

	statusArray := make([]string, 0, len(obj.ResultData.InfoData.Status))
	for _, status := range obj.ResultData.InfoData.Status {
		statusArray = append(statusArray, string(status.DomainStatusType))
	}
	buffer.WriteString(fmt.Sprintf("status[%s]", strings.Join(statusArray, ",")))

	ctx.String(200, buffer.String())
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
		layoutFormat := "2006-01-02T15:04:05.999999999Z"

		newExpireDate, err := time.Parse(layoutFormat, obj.ResultData.RenewedData.ExpiredDate)
		if err != nil {
			err = errors.Wrap(err, "DomainController Renew: time.Parse")
		}

		newExpireDate = newExpireDate.Local()
		newExpireDate = newExpireDate.Add(-(time.Hour * 8))

		res = fmt.Sprintf("%v %v", obj.Result.Code, newExpireDate.Format("2006-01-02 15:04:05"))
	} else {
		res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)
	}

	ctx.String(200, res)
	return
}
