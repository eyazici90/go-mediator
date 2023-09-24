package mediator

type errType int

const (
	ErrHandlerNotFound errType = iota
	ErrInvalidArg

	errCount
)

func (e errType) Error() string {
	if e < 0 || e >= errCount {
		panic("invalid err number")
	}
	return errDescriptions[e]
}

var _ = [1]int{}[len(errDescriptions)-int(errCount)]

var errDescriptions = [...]string{
	ErrHandlerNotFound: "handler could not be found",
	ErrInvalidArg:      "invalid arg",
}
