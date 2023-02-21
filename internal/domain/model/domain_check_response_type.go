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
	Name   Name   `xml:"name"`
	Reason string `xml:"reason,omitempty"`
}

type Name struct {
	AvailKey int    `xml:"avail,attr"`
	Value    string `xml:",chardata"`
}
