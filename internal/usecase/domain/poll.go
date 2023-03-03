package domain

type PollRequestResponseDTO interface {
	GetMessageQueueId() *string
	GetResultCode() int
	GetMessageQueue() any
}
