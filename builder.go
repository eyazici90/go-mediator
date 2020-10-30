package mediator

import (
	"context"
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

func (b *builder) Use(call func(context.Context, Message, Next) error) Builder {
	b.pipelineContext.Behaviours = append(b.pipelineContext.Behaviours, call)
	return b
}

func (b *builder) RegisterHandler(request Message, handler RequestHandler) Builder {
	key := request.Key()

	b.handlers[key] = handler
	return b
}

func (b *builder) Build() (Mediator, error) {
	m := newMediator(b.pipelineContext, b.handlers)
	reverseApply(b.pipelineContext.Behaviours, m.pipe)
	return m, nil
}

func reverseApply(behaviours []Behaviour,
	action func(Behaviour)) {
	for i := len(behaviours) - 1; i >= 0; i-- {
		action(behaviours[i])
	}
}
