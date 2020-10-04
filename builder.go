package mediator

import (
	"context"
	"reflect"
)

type builder struct {
	pipelineContext PipelineContext
	handlers        map[string]RequestHandler
}

func New() Builder {
	return &builder{
		pipelineContext: NewPipelineContext(),
		handlers:        make(map[string]RequestHandler),
	}
}

func (b *builder) UseBehaviour(pipelineBehaviour PipelineBehaviour) Builder {
	return b.Use(pipelineBehaviour.Process)
}

func (b *builder) Use(call func(context.Context, interface{}, Next) error) Builder {
	b.pipelineContext.Behaviours = append(b.pipelineContext.Behaviours, call)
	return b
}

func (b *builder) RegisterHandler(request interface{}, handler RequestHandler) Builder {
	requestType := reflect.TypeOf(request).Name()

	b.handlers[requestType] = handler
	return b
}

func (b *builder) Build() (Mediator, error) {
	m := newMediator(b.pipelineContext, b.handlers)
	ReverseApply(b.pipelineContext.Behaviours, m.pipe)
	return m, nil
}
