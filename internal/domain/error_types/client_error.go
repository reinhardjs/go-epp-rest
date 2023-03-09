package error_types

type InteractorError struct {
	Original error
}

func (e *InteractorError) Error() string {
	return e.Original.Error()
}

type ControllerError struct {
	Original error
}

func (e *ControllerError) Error() string {
	return e.Original.Error()
}

type PresenterError struct {
	Original error
}

func (e *PresenterError) Error() string {
	return e.Original.Error()
}
