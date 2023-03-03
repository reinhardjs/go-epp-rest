package mapper

import (
	"encoding/xml"

	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter"
)

type XMLMapperImpl struct{}

func NewXMLMapper() adapter.XMLMapper {
	return &XMLMapperImpl{}
}

func (m *XMLMapperImpl) Decode(origin []byte, destination interface{}) error {
	err := xml.Unmarshal(origin, &destination)
	return err
}
