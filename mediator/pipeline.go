package mediator

import "context"

type (
	Next func(ctx context.Context) error
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

type pipeline struct {
	bhs      behaviors
	handlers []Handler
}

func (p *pipeline) useBehavior(behavior PipelineBehaviour) error {
	if behavior == nil {
		return ErrInvalidArg
	}
	return p.use(behavior.Process)
}

func (p *pipeline) use(call func(context.Context, Message, Next) error) error {
	if call == nil {
		return ErrInvalidArg
	}
	p.bhs = append(p.bhs, call)
	return nil
}

func (p *pipeline) registerHandler(req Message, h Handler) error {
	if req == nil || h == nil {
		return ErrInvalidArg
	}
	key := req.Key()
	p.handlers[key] = h

	return nil
}

func (p *pipeline) findHandler(key int) (Handler, error) {
	v := p.handlers[key]
	if v == nil {
		return nil, ErrHandlerNotFound
	}
	return v, nil
}
