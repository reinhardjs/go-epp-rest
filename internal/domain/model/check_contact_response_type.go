package model

// Response represents an EPP response.
type CheckContactResponse struct {
	Result        Result                 `xml:"response>result"`
	ResultData    CheckContactResultData `xml:"response>resData"`
	TransactionID TransactionID          `xml:"response>trID"`
}

type CheckContactResultData struct {
	CheckDatas []CheckContactData `xml:"chkData>cd"`
}

type CheckContactData struct {
	Id     Id     `xml:"id"`
	Reason string `xml:"reason,omitempty"`
}

type Id struct {
	AvailKey int    `xml:"avail,attr"`
	Value    string `xml:",chardata"`
}
