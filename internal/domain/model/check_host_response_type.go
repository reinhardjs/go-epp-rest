package model

// Response represents an EPP response.
type CheckHostResponse struct {
	Result        []Result            `xml:"response>result"`
	ResultData    CheckHostResultData `xml:"response>resData"`
	TransactionID TransactionID       `xml:"response>trID"`
}

type CheckHostResultData struct {
	CheckDatas []CheckHostData `xml:"chkData>cd"`
}

type CheckHostData struct {
	HostName HostName `xml:"name"`
	Reason   string   `xml:"reason,omitempty"`
}

type HostName struct {
	AvailKey int    `xml:"avail,attr"`
	Value    string `xml:",chardata"`
}
