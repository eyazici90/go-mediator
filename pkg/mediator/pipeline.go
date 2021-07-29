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

type Option func(pCtx *PipelineContext)

func (p Pipeline) empty() bool { return p == nil }

type PipelineContext struct {
	behaviors Behaviors
	pipeline  Pipeline
	handlers  map[string]RequestHandler
}

func NewContext(opts ...Option) *PipelineContext {
	ctx := &PipelineContext{
		handlers: make(map[string]RequestHandler),
	}
	for _, opt := range opts {
		opt(ctx)
	}
	return ctx
}

func WithBehaviour(behavior PipelineBehaviour) Option {
	return func(pCtx *PipelineContext) {
		pCtx.useBehavior(behavior)
	}
}

func WithBehaviourFunc(fn func(context.Context, Message, Next) error) Option {
	return func(pCtx *PipelineContext) {
		pCtx.use(fn)
	}
}

func WithHandler(req Message, rh RequestHandler) Option {
	return func(pCtx *PipelineContext) {
		pCtx.registerHandler(req, rh)
	}
}

func (p *PipelineContext) useBehavior(behavior PipelineBehaviour) {
	p.use(behavior.Process)
}

func (p *PipelineContext) use(call func(context.Context, Message, Next) error) {
	p.behaviors = append(p.behaviors, call)
}

func (p *PipelineContext) registerHandler(req Message, h RequestHandler) {
	key := req.Key()
	p.handlers[key] = h
}

func (p *PipelineContext) Build() (*Mediator, error) {
	m := newMediator(*p)
	p.behaviors.reverseApply(m.pipe)
	return m, nil
}
