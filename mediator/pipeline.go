package mediator

import "context"

type Behaviour func(context.Context, Message, Next) error

type Next func(ctx context.Context) error

type Pipeline func(context.Context, Message) error

func (p Pipeline) empty() bool { return p == nil }

type PipelineContext struct {
	behaviours []Behaviour
	pipeline   Pipeline
	handlers   map[string]RequestHandler
}

func NewContext() *PipelineContext {
	return &PipelineContext{
		handlers: make(map[string]RequestHandler),
	}
}

func (p *PipelineContext) UseBehaviour(behaviour PipelineBehaviour) Builder {
	return p.Use(behaviour.Process)
}

func (p *PipelineContext) Use(call func(context.Context, Message, Next) error) Builder {
	p.behaviours = append(p.behaviours, call)
	return p
}

func (p *PipelineContext) RegisterHandler(req Message, h RequestHandler) Builder {
	key := req.Key()

	p.handlers[key] = h
	return p
}

func (p *PipelineContext) Build() (*Mediator, error) {
	m := newMediator(*p)
	reverseApply(p.behaviours, m.pipe)
	return m, nil
}

func reverseApply(behaviours []Behaviour,
	fn func(Behaviour)) {
	for i := len(behaviours) - 1; i >= 0; i-- {
		fn(behaviours[i])
	}
}
