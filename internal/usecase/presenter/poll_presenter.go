package presenter

type PollPresenter interface {
	Poll(responseObject PollRequestResponseDTO) string
}

type PollRequestResponseDTO interface {
	GetResultCode() int
}
