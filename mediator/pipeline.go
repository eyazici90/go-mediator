package mediator

import "context"

type (
	Option func(pipe *Pipeline) error
	Next   func(ctx context.Context) error
)

type (
	behavior  func(context.Context, Message, Next) error
	behaviors []behavior
)

func (b behaviors) merge() behavior {
	var result func(context.Context, Message, Next) error

	for _, v := range b {
		v := v
		if result == nil {
			result = v
			continue
		}
		seed := result
		result = func(ctx context.Context, msg Message, next Next) error {
			return seed(ctx, msg, func(ctx context.Context) error {
				return v(ctx, msg, next)
			})
		}
	}

	return result
}

const maxSize = 64

type PipelineFunc func(context.Context, Message) error

type Pipeline struct {
	bhs      behaviors
	handlers []Handler
}

func newPipeline(opts ...Option) (*Pipeline, error) {
	pipe := Pipeline{
		handlers: make([]Handler, maxSize),
	}
	for _, opt := range opts {
		if err := opt(&pipe); err != nil {
			return nil, err
		}
	}
	return &pipe, nil
}

func WithBehaviour(behavior PipelineBehaviour) Option {
	return func(pipe *Pipeline) error {
		return pipe.useBehavior(behavior)
	}
}

func WithBehaviourFunc(fn func(context.Context, Message, Next) error) Option {
	return func(pipe *Pipeline) error {
		return pipe.use(fn)
	}
}

func WithHandler(req Message, rh Handler) Option {
	return func(pipe *Pipeline) error {
		return pipe.registerHandler(req, rh)
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
	p.bhs = append(p.bhs, call)
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
