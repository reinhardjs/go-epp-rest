package model

// Response represents an EPP response.
type DomainCheckResponse struct {
	Result        []Result              `xml:"response>result"`
	ResultData    DomainCheckResultData `xml:"response>resData"`
	TransactionID TransactionID         `xml:"response>trID"`
}

type DomainCheckResultData struct {
	CheckDatas []DomainCheckData `xml:"chkData>cd"`
}

type DomainCheckData struct {
	Name   string `xml:"name"`
	Reason string `xml:"reason,omitempty"`
}
