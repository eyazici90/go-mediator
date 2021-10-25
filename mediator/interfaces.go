package mediator

import "context"

type (
	Sender interface {
		Send(context.Context, Message) error
	}
	RequestHandler interface {
		Handle(context.Context, Message) error
	}
	PipelineBehaviour interface {
		Process(context.Context, Message, Next) error
	}
	Message interface {
		Key() string
	}
)
