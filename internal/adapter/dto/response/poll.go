package response

import "gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"

// Poll Ack
type PollAckResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}

// Poll Request
type PollRequestResponse struct {
	Result        Result          `xml:"response>result"`
	MessageQueue  *MessageQueue   `xml:"response>msgQ"`
	ResultData    *PollResultData `xml:"response>resData"`
	TransactionID TransactionID   `xml:"response>trID"`
}

type PollResultData struct {
	InfoData     *types.DomainInfoData     `xml:"urn:ietf:params:xml:ns:domain-1.0 infData"`
	CreateData   *types.DomainCreateData   `xml:"urn:ietf:params:xml:ns:domain-1.0 creData"`
	TransferData *types.DomainTransferData `xml:"urn:ietf:params:xml:ns:domain-1.0 trnData"`
	RenewData    *types.DomainRenewData    `xml:"urn:ietf:params:xml:ns:domain-1.0 renData"`
}

type MessageQueue struct {
	Count     string `xml:"count,attr"`
	Id        string `xml:"id,attr"`
	QueueDate string `xml:"qDate"`
	Message   string `xml:"msg"`
}
