package presenter

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"gitlab.com/merekmu/go-epp-rest/internal/domain/model"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/presenter"
)

type domainPresenter struct{}

func NewDomainPresenter() presenter.RegistrarPresenter {
	return &domainPresenter{}
}

func (p *domainPresenter) ResponseCheck(response []byte) (responseString string, err error) {
	responseObj := model.DomainCheckResponse{}

	if err := xml.Unmarshal(response, &responseObj); err != nil {
		log.Println(errors.Wrap(err, "Domain Controller: CheckDomain xml.Unmarshal"))
	}

	for _, element := range responseObj.ResultData.CheckDatas {
		notStr := ""
		if element.Name.AvailKey == 0 {
			notStr = "not "
		}
		responseString += fmt.Sprintf("Domain %s, domain %savailable\n", element.Name.Value, notStr)
	}

	responseString = strings.TrimSuffix(responseString, "\n")

	return
}
