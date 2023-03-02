package mapper

import (
	"encoding/xml"

	"gitlab.com/merekmu/go-epp-rest/internal/usecase/adapter"
)

type XMLMapperImpl struct{}

func NewXMLMapper() adapter.XMLMapper {
	return &XMLMapperImpl{}
}

func (m *XMLMapperImpl) MapXMLToModel(origin []byte, destination interface{}) error {
	err := xml.Unmarshal(origin, &destination)
	return err
}
