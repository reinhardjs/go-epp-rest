package error_types

type RequestTimeOutError struct {
	Detail string
}

func (e *RequestTimeOutError) Error() string {
	return "request time out: " + e.Detail
}
