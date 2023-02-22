package model

// Response represents an EPP response.
type DeleteDomainResponse struct {
	Result        []Result      `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}
