package mediator

import "context"

type (
	Mediator interface {
		sender
		publisher
	}
	sender interface {
		Send(context.Context, interface{}) error
	}
	publisher interface {
		Publish(msg interface{})
	}
	Builder interface {
		RegisterHandler(request interface{}, handler RequestHandler) Builder
		UseBehaviour(PipelineBehaviour) Builder
		Use(call func(context.Context, interface{}, Next) error) Builder
		Build() (Mediator, error)
	}
)
