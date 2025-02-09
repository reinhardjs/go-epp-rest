package response

type Response struct {
	Result Result `xml:"response>result"`
}

// TransactionID represents transaction IDs for the client and the server.
type TransactionID struct {
	ClientTransactionID string `xml:"clTRID,omitempty"`
	ServerTransactionID string `xml:"svTRID"`
}

// Result represents the result in a EPP response.
type Result struct {
	Code          int                 `xml:"code,attr"`
	Message       string              `xml:"msg"`
	Value         Value               `xml:"value"`
	ExternalValue *ExternalErrorValue `xml:"extValue,omitempty"`
}

type Value struct {
	Texts []string `xml:"text"`
}

// ExternalErrorValue represents the response in the extValeu tag.
type ExternalErrorValue struct {
	Value  ExtValue `xml:"value"`
	Reason string   `xml:"reason"`
}

type ExtValue struct {
	ReasonCode string `xml:"reasonCode"`
}
