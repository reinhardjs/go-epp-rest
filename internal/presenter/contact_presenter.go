package presenter

import (
	"bytes"
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

func (p *contactPresenter) CheckSuccess(ctx infrastructure.Context, obj response.CheckContactResponse) (err error) {
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
	return
}

func (p *contactPresenter) CreateSuccess(ctx infrastructure.Context, obj response.CreateContactResponse) (err error) {
	var res string
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("1000 %s", obj.ResultData.CreateData.Id))

	res = buffer.String()
	ctx.String(200, res)
	return
}

func (p *contactPresenter) UpdateSuccess(ctx infrastructure.Context, obj response.UpdateContactResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
	return
}

func (p *contactPresenter) DeleteSuccess(ctx infrastructure.Context, obj response.DeleteContactResponse) (err error) {
	var res string

	res = fmt.Sprintf("%v %v", obj.Result.Code, obj.Result.Message)

	ctx.String(200, res)
	return
}

func (p *contactPresenter) InfoSuccess(ctx infrastructure.Context, obj response.InfoContactResponse) (err error) {
	var res string
	var buffer bytes.Buffer

	buffer.WriteString("1000 CONTACT INFO IS::\n")
	buffer.WriteString(fmt.Sprintf("ID :%s\n", obj.ResultData.InfoData.Name))
	buffer.WriteString(fmt.Sprintf("ROID :%s\n", obj.ResultData.InfoData.ROID))
	for _, postalInfo := range obj.ResultData.InfoData.PostalInfo {
		buffer.WriteString(fmt.Sprintf("NAME :%s\n", postalInfo.Name))
		buffer.WriteString(fmt.Sprintf("ORG :%s\n", postalInfo.Organization))
		for index, address := range postalInfo.Address.Street {
			buffer.WriteString(fmt.Sprintf("STREET%d :%s\n", index, address))
		}
		buffer.WriteString(fmt.Sprintf("CITY :%s\n", postalInfo.Address.City))
		buffer.WriteString(fmt.Sprintf("SP :%s\n", postalInfo.Address.StateProvince))
		buffer.WriteString(fmt.Sprintf("PC :%s\n", postalInfo.Address.PostalCode))
		buffer.WriteString(fmt.Sprintf("CC :%s\n", postalInfo.Address.CountryCode))
	}
	buffer.WriteString(fmt.Sprintf("VOICE :%s\n", obj.ResultData.InfoData.Voice.Value))
	buffer.WriteString(fmt.Sprintf("FAX :%s\n", obj.ResultData.InfoData.Fax.Value))
	buffer.WriteString(fmt.Sprintf("EMAIL :%s\n", obj.ResultData.InfoData.Email))
	buffer.WriteString(fmt.Sprintf("clID :%s\n", obj.ResultData.InfoData.ClientID))
	buffer.WriteString(fmt.Sprintf("crID :%s\n", obj.ResultData.InfoData.CreateID))
	buffer.WriteString(fmt.Sprintf("crDate :%s\n", obj.ResultData.InfoData.CreateDate.Local().Format("2006-01-02 15:04:05")))
	buffer.WriteString(fmt.Sprintf("upID :%s\n", obj.ResultData.InfoData.UpdateID))
	buffer.WriteString(fmt.Sprintf("authInfo :%s", obj.ResultData.InfoData.AuthInfo.Password))

	res = buffer.String()

	ctx.String(200, res)
	return
}
