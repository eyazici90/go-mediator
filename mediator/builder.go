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

func (b *builder) UseBehaviour(p PipelineBehaviour) Builder {
	return b.Use(p.Process)
}

func (b *builder) Use(call func(context.Context, Message, Next) error) Builder {
	b.pipelineContext.Behaviours = append(b.pipelineContext.Behaviours, call)
	return b
}

func (b *builder) RegisterHandler(req Message, h RequestHandler) Builder {
	key := req.Key()

	b.handlers[key] = h
	return b
}

func (b *builder) Build() (Mediator, error) {
	m := newMediator(b.pipelineContext, b.handlers)
	reverseApply(b.pipelineContext.Behaviours, m.pipe)
	return m, nil
}

func reverseApply(behaviours []Behaviour,
	fn func(Behaviour)) {
	for i := len(behaviours) - 1; i >= 0; i-- {
		fn(behaviours[i])
	}
}
