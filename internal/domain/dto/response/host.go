package response

import "gitlab.com/merekmu/go-epp-rest/pkg/registry_epp/types"

// Host Check
type CheckHostResponse struct {
	Result        Result              `xml:"response>result"`
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
	Available int    `xml:"avail,attr"`
	Value     string `xml:",chardata"`
}

// Host Create
type CreateHostResponse struct {
	Result        Result               `xml:"response>result"`
	ResultData    CreateHostResultData `xml:"response>resData"`
	TransactionID TransactionID        `xml:"response>trID"`
}

type CreateHostResultData struct {
	CreateData CreateHostData `xml:"creData"`
}

type CreateHostData struct {
	Name       string `xml:"name"`
	CreateDate string `xml:"crDate"`
}

// Host Update
type UpdateHostResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}

// Host Delete
type DeleteHostResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}

// Host Info
type InfoHostResponse struct {
	Result        Result             `xml:"response>result"`
	ResultData    InfoHostResultData `xml:"response>resData"`
	TransactionID TransactionID      `xml:"response>trID"`
}

type InfoHostResultData struct {
	InfoData types.HostInfoData `xml:"urn:ietf:params:xml:ns:host-1.0 infData"`
}
