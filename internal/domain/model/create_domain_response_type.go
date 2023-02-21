package model

// Response represents an EPP response.
type CreateDomainResponse struct {
	Result        []Result               `xml:"response>result"`
	ResultData    CreateDomainResultData `xml:"response>resData"`
	TransactionID TransactionID          `xml:"response>trID"`
}

type CreateDomainResultData struct {
	CheckDatas []CreateDomainData `xml:"chkData>cd"`
}

type CreateDomainData struct {
	Name   Name   `xml:"name"`
	Reason string `xml:"reason,omitempty"`
}

type DomainName struct {
	AvailKey int    `xml:"avail,attr"`
	Value    string `xml:",chardata"`
}
