package model

// Response represents an EPP response.
type CreateDomainResponse struct {
	Result        Result                 `xml:"response>result"`
	ResultData    CreateDomainResultData `xml:"response>resData"`
	TransactionID TransactionID          `xml:"response>trID"`
}

type CreateDomainResultData struct {
	CreatedData CreateDomainData `xml:"creData>cd"`
}

type CreateDomainData struct {
	Name        string `xml:"name"`
	CreatedDate string `xml:"crDate"`
	ExpiredDate string `xml:"exDate"`
}
