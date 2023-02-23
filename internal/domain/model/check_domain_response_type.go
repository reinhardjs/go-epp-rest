package model

// Response represents an EPP response.
type CheckDomainResponse struct {
	Result        Result                `xml:"response>result"`
	ResultData    CheckDomainResultData `xml:"response>resData"`
	TransactionID TransactionID         `xml:"response>trID"`
}

type CheckDomainResultData struct {
	CheckDatas []CheckDomainData `xml:"chkData>cd"`
}

type CheckDomainData struct {
	Name   Name   `xml:"name"`
	Reason string `xml:"reason,omitempty"`
}

type Name struct {
	AvailKey int    `xml:"avail,attr"`
	Value    string `xml:",chardata"`
}
