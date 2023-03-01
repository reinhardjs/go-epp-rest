package mapper

import (
	"encoding/xml"

	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/usecase/infrastructure"
)

type XMLMapperImpl struct{}

func NewXMLMapper() infrastructure.XMLMapper {
	return &XMLMapperImpl{}
}

func (m *XMLMapperImpl) MapXMLToModel(input string) (res response.PollRequestResponse, err error) {
	err = xml.Unmarshal([]byte(input), &res)
	return
}
