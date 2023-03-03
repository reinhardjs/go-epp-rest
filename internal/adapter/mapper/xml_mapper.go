package mapper

import (
	"encoding/xml"

	"gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"
	"gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter/mapper"
)

type XMLMapperImpl struct{}

func NewXMLMapper() mapper.XMLMapper {
	return &XMLMapperImpl{}
}

func (m *XMLMapperImpl) Decode(origin []byte, destination interface{}) error {
	err := xml.Unmarshal(origin, &destination)
	return err
}

func (m *XMLMapperImpl) ToPollRequestResponseDTO(origin []byte) (output *response.PollRequestResponse, err error) {
	err = m.Decode(origin, &output)
	return
}
