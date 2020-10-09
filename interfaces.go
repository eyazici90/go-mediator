package mediator

import "context"

type (
	Mediator interface {
		Sender
		Publisher
	}
	Sender interface {
		Send(context.Context, interface{}) error
	}
	Publisher interface {
		Publish(msg interface{})
	}
	Builder interface {
		RegisterHandler(request interface{}, handler RequestHandler) Builder
		UseBehaviour(PipelineBehaviour) Builder
		Use(call func(context.Context, interface{}, Next) error) Builder
		Build() (Mediator, error)
	}
)
