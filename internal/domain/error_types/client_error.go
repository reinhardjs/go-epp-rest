package error_types

type InteractorError struct {
	Message  string
	Original error
}

func (e *InteractorError) Error() string {
	if e.Original != nil {
		return e.Original.Error()
	}

	return e.Message
}

type ControllerError struct {
	Message  string
	Original error
}

func (e *ControllerError) Error() string {
	if e.Original != nil {
		return e.Original.Error()
	}

	return e.Message
}

type PresenterError struct {
	Message  string
	Original error
}

func (e *PresenterError) Error() string {
	if e.Original != nil {
		return e.Original.Error()
	}

	return e.Message
}
