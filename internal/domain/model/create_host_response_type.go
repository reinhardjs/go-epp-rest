package model

// Response represents an EPP response.
type CreateHostResponse struct {
	Result        []Result             `xml:"response>result"`
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
