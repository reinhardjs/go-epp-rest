package infrastructure

import "gitlab.com/merekmu/go-epp-rest/internal/interfaces/adapter/dto/response"

type XMLMapper interface {
	MapXMLToModel(xml string) (response.PollRequestResponse, error)
}
