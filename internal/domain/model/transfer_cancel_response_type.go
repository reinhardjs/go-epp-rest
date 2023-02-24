package model

// Response represents an EPP response.
type TransferCancelResponse struct {
	Result        Result        `xml:"response>result"`
	TransactionID TransactionID `xml:"response>trID"`
}
