package response

// provide implementations to be used by external layer via interface

// -- PollRequestResponse --
func (r *PollRequestResponse) GetMessageQueueId() *string {
	return &r.MessageQueue.Id
}

func (r *PollRequestResponse) GetResultCode() int {
	return r.Result.Code
}

func (r *PollRequestResponse) GetMessageQueue() any {
	return r.MessageQueue
}
