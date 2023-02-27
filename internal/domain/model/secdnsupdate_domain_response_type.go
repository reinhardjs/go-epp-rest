package model

import "gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"

// Response represents an EPP response.
type SecDNSUpdateResponse struct {
	Result        Result                 `xml:"response>result"`
	ResultData    SecDNSUpdateResultData `xml:"response>resData"`
	TransactionID TransactionID          `xml:"response>trID"`
}

type SecDNSUpdateResultData struct {
	types.DomainInfoData `xml:"urn:ietf:params:xml:ns:domain-1.0 infData"`
}
