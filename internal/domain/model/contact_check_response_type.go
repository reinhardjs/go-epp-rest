package model

// Response represents an EPP response.
type ContactCheckResponse struct {
	Result        []Result               `xml:"response>result"`
	ResultData    ContactCheckResultData `xml:"response>resData"`
	TransactionID TransactionID          `xml:"response>trID"`
}

type ContactCheckResultData struct {
	CheckDatas []ContactCheckData `xml:"chkData>cd"`
}

type ContactCheckData struct {
	Id     Id     `xml:"id"`
	Reason string `xml:"reason,omitempty"`
}

type Id struct {
	AvailKey int    `xml:"avail,attr"`
	Value    string `xml:",chardata"`
}
