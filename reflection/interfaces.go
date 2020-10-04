package reflection

import (
	"context"

	"github.com/eyazici90/go-mediator"
)

type (
	Mediator interface {
		initializer
		sender
		publisher
		pipelineBuilder
	}
	initializer interface {
		RegisterHandler(handler interface{}) Mediator
		RegisterHandlers(handlers ...interface{}) Mediator
	}
	sender interface {
		Send(context.Context, interface{}) error
	}
	pipelineBuilder interface {
		Build() Mediator
		UseBehaviour(mediator.PipelineBehaviour) Mediator
		Use(call func(context.Context, interface{}, mediator.Next) error) Mediator
	}
	publisher interface {
		Publish(msg interface{})
	}
)
