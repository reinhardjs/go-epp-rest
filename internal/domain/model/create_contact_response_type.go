package model

// Response represents an EPP response.
type CreateDomainResponse struct {
	Result        []Result               `xml:"response>result"`
	ResultData    CreateDomainResultData `xml:"response>resData"`
	TransactionID TransactionID          `xml:"response>trID"`
}

type CreateDomainResultData struct {
	CreatedData CreateDomainData `xml:"chkData>cd"`
}

type CreateDomainData struct {
	Id          string `xml:"name"`
	CreatedDate string `xml:"reason,omitempty"`
}
