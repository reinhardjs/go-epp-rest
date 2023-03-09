package error_types

type HostNameNotAvailableError struct{}

func (e *HostNameNotAvailableError) Error() string {
	return "host name not available"
}
