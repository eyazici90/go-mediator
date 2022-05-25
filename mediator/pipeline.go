package mediator

import "context"

type (
	Option func(pCtx *Pipeline) error
	Next   func(ctx context.Context) error
)

type (
	Behavior  func(context.Context, Message, Next) error
	Behaviors []Behavior
)

func (b Behaviors) merge() Behavior {
	var result func(context.Context, Message, Next) error

	for _, v := range b {
		if result == nil {
			result = v
			continue
		}
		seed := result
		result = func(ctx context.Context, msg Message, next Next) error {
			return v(ctx, msg, func(ctx context.Context) error {
				return seed(ctx, msg, next)
			})
		}
	}

	return result
}

const maxSize = 64

type PipelineFunc func(context.Context, Message) error

type Pipeline struct {
	call PipelineFunc

	behaviors Behaviors
	handlers  []Handler
}

func newPipeline(opts ...Option) (*Pipeline, error) {
	ctx := Pipeline{
		handlers: make([]Handler, maxSize),
	}
	for _, opt := range opts {
		if err := opt(&ctx); err != nil {
			return nil, err
		}
	}
	return &ctx, nil
}

func WithBehaviour(behavior PipelineBehaviour) Option {
	return func(pCtx *Pipeline) error {
		return pCtx.useBehavior(behavior)
	}
}

func WithBehaviourFunc(fn func(context.Context, Message, Next) error) Option {
	return func(pCtx *Pipeline) error {
		return pCtx.use(fn)
	}
}

func WithHandler(req Message, rh Handler) Option {
	return func(pCtx *Pipeline) error {
		return pCtx.registerHandler(req, rh)
	}
}

func (p *Pipeline) useBehavior(behavior PipelineBehaviour) error {
	if behavior == nil {
		return ErrInvalidArg
	}
	return p.use(behavior.Process)
}

func (p *Pipeline) use(call func(context.Context, Message, Next) error) error {
	if call == nil {
		return ErrInvalidArg
	}
	p.behaviors = append(p.behaviors, call)
	return nil
}

func (p *Pipeline) registerHandler(req Message, h Handler) error {
	if req == nil || h == nil {
		return ErrInvalidArg
	}
	key := req.Key()
	p.handlers[key] = h

	return nil
}

func (p *Pipeline) findHandler(key int) (Handler, error) {
	v := p.handlers[key]
	if v == nil {
		return nil, ErrHandlerNotFound
	}
	return v, nil
}
