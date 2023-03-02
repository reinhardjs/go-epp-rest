package presenter

type PollPresenter interface {
	Acknowledge(responseObject PollResponse) string
	Request(responseObject PollResponse) string
}

type PollResponse interface {
	GetResultCode() int
	GetResultMessage() string
}
