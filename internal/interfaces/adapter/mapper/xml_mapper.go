package mapper

import (
	"encoding/xml"

	"gitlab.com/merekmu/go-epp-rest/internal/usecase/infrastructure"
)

type XMLMapperImpl struct{}

func NewXMLMapper() infrastructure.XMLMapper {
	return &XMLMapperImpl{}
}

func (m *XMLMapperImpl) MapXMLToModel(input string, model interface{}) error {
	return xml.Unmarshal([]byte(input), model)
}
