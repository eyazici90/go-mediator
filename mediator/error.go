package mediator

import "errors"

var (
	ErrHandlerNotFound = errors.New("handler could not be found")
	ErrInvalidArg      = errors.New("invalid arg")
)
