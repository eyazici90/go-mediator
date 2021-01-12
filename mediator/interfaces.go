package mediator

import "context"

type (
	Mediator interface {
		Sender
		Publisher
	}
	Sender interface {
		Send(context.Context, Message) error
	}
	Publisher interface {
		Publish(msg Message)
	}
	Builder interface {
		RegisterHandler(request Message, handler RequestHandler) Builder
		UseBehaviour(PipelineBehaviour) Builder
		Use(fn func(context.Context, Message, Next) error) Builder
		Build() (Mediator, error)
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
