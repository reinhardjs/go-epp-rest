package model

import "gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"

// Response represents an EPP response.
type TransferCheckResponse struct {
	Result        Result               `xml:"response>result"`
	ResultData    DomainInfoResultData `xml:"response>resData"`
	TransactionID TransactionID        `xml:"response>trID"`
}

type TransferCheckResultData struct {
	CheckData types.TransferData `xml:"urn:ietf:params:xml:ns:domain-1.0 Data"`
}
