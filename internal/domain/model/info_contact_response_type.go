package model

import (
	"gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"
)

// Response represents an EPP response.
type InfoContactResponse struct {
	Result        Result               `xml:"response>result"`
	ResultData    DomainInfoResultData `xml:"response>resData"`
	TransactionID TransactionID        `xml:"response>trID"`
}

type DomainContactResultData struct {
	InfoData types.ContactInfoData `xml:"urn:ietf:params:xml:ns:domain-1.0 infData"`
}
