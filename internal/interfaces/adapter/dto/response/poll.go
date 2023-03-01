package response

// Poll Ack
type PollAckResponse struct {
	Result        Result        `xml:"response>result"`
	MessageQueue  MessageQueue  `xml:"response>msgQ"`
	TransactionID TransactionID `xml:"response>trID"`
}

// Poll Request
type PollRequestResponse struct {
	Result        Result               `xml:"response>result"`
	MessageQueue  MessageQueue         `xml:"response>msgQ"`
	ResultData    DomainInfoResultData `xml:"response>resData"`
	TransactionID TransactionID        `xml:"response>trID"`
}

type MessageQueue struct {
	Count     string `xml:"count,attr"`
	Id        string `xml:"id,attr"`
	QueueDate string `xml:"qDate"`
	Message   string `xml:"msg"`
}
