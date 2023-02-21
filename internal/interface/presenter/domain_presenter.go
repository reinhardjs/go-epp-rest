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

type domainPresenter[T model.DomainCheckResponse] struct{}

func NewDomainPresenter() presenter.RegistrarPresenter[model.DomainCheckResponse] {
	return &domainPresenter[model.DomainCheckResponse]{}
}

func (p *domainPresenter[T]) ResponseCheck(responseString string) (response string, err error) {
	responseObj := model.DomainCheckResponse{}

	log.Println(responseString)

	if err := xml.Unmarshal([]byte(responseString), &responseObj); err != nil {
		log.Println(errors.Wrap(err, "Domain Controller: CheckDomain xml.Unmarshal"))
	}

	for _, element := range responseObj.ResultData.CheckDatas {
		notStr := ""
		if element.Name.AvailKey == 0 {
			notStr = "not "
		}
		response += fmt.Sprintf("Domain %s, domain %savailable\n", element.Name.Value, notStr)
	}

	response = strings.TrimSuffix(response, "\n")

	return
}
