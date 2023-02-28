package responses

import "gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"

// Transfer Check
type TransferCheckResponse struct {
	Result        Result               `xml:"response>result"`
	ResultData    DomainInfoResultData `xml:"response>resData"`
	TransactionID TransactionID        `xml:"response>trID"`
}

type TransferCheckResultData struct {
	CheckData types.TransferData `xml:"urn:ietf:params:xml:ns:domain-1.0 Data"`
}

// Transfer Request
type TransferRequestResponse struct {
	Result        Result               `xml:"response>result"`
	ResultData    DomainInfoResultData `xml:"response>resData"`
	TransactionID TransactionID        `xml:"response>trID"`
}

type TransferRequestResultData struct {
	TransferData types.TransferData `xml:"urn:ietf:params:xml:ns:domain-1.0 trnData"`
}

// Transfer Approve
type TransferApproveResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}

// Transfer Reject
type TransferRejectResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}

// Transfer Cancel
type TransferCancelResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}
