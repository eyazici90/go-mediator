package mediator

import "context"

type (
	Mediator interface {
		initializer
		sender
		publisher
		pipelineBuilder
	}

	initializer interface {
		RegisterHandler(handler RequestHandler) Mediator
		RegisterHandlers(handlers ...RequestHandler) Mediator
	}
	sender interface {
		Send(context.Context, interface{}) error
	}
	pipelineBuilder interface {
		Build() Mediator
		UseBehaviour(PipelineBehaviour) Mediator
		Use(call func(context.Context, interface{}, Next) error) Mediator
	}
	publisher interface {
		Publish(msg interface{})
	}
)
