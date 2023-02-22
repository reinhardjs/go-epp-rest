package model

// Response represents an EPP response.
type CreateContactResponse struct {
	Result        []Result                `xml:"response>result"`
	ResultData    CreateContactResultData `xml:"response>resData"`
	TransactionID TransactionID           `xml:"response>trID"`
}

type CreateContactResultData struct {
	CreateData CreateContactData `xml:"creData"`
}

type CreateContactData struct {
	Id         string `xml:"id"`
	CreateDate string `xml:"crDate"`
}
