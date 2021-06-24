package mediator

import "context"

type (
	Behavior  func(context.Context, Message, Next) error
	Behaviors []Behavior
)

func (b Behaviors) reverseApply(fn func(Behavior)) {
	for i := len(b) - 1; i >= 0; i-- {
		fn(b[i])
	}
}

type Next func(ctx context.Context) error

type Pipeline func(context.Context, Message) error

func (p Pipeline) empty() bool { return p == nil }

type PipelineContext struct {
	behaviors []Behavior
	pipeline  Pipeline
	handlers  map[string]RequestHandler
}

func NewContext() *PipelineContext {
	return &PipelineContext{
		handlers: make(map[string]RequestHandler),
	}
}

func (p *PipelineContext) UseBehavior(behavior PipelineBehaviour) Builder {
	return p.Use(behavior.Process)
}

func (p *PipelineContext) Use(call func(context.Context, Message, Next) error) Builder {
	p.behaviors = append(p.behaviors, call)
	return p
}

func (p *PipelineContext) RegisterHandler(req Message, h RequestHandler) Builder {
	key := req.Key()

	p.handlers[key] = h
	return p
}

func (p *PipelineContext) Build() (*Mediator, error) {
	m := newMediator(*p)
	Behaviors(p.behaviors).reverseApply(m.pipe)
	return m, nil
}
