package mediator

import "context"

type (
	Handler interface {
		Handle(context.Context, Message) error
	}
	PipelineBehaviour interface {
		Process(context.Context, Message, Next) error
	}
	Message interface {
		Key() int
	}
)
