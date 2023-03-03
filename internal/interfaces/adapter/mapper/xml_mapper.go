package mapper

import "gitlab.com/merekmu/go-epp-rest/internal/domain/dto/response"

type XMLMapper interface {
	Decode(origin []byte, destination interface{}) error
	ToPollRequestResponseDTO(origin []byte) (output *response.PollRequestResponse, err error)
}
